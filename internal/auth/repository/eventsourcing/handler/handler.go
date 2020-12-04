package handler

import (
	"time"

	"github.com/caos/zitadel/internal/auth/repository/eventsourcing/view"
	sd "github.com/caos/zitadel/internal/config/systemdefaults"
	"github.com/caos/zitadel/internal/config/types"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/spooler"
	iam_events "github.com/caos/zitadel/internal/iam/repository/eventsourcing"
	key_view_model "github.com/caos/zitadel/internal/key/repository/view/model"
	org_events "github.com/caos/zitadel/internal/org/repository/eventsourcing"
	proj_event "github.com/caos/zitadel/internal/project/repository/eventsourcing"
	usr_event "github.com/caos/zitadel/internal/user/repository/eventsourcing"
	user_view_model "github.com/caos/zitadel/internal/user/repository/view/model"
)

type Configs map[string]*Config

type Config struct {
	MinimumCycleDuration types.Duration
}

type handler struct {
	view                *view.View
	bulkLimit           uint64
	cycleDuration       time.Duration
	errorCountUntilSkip uint64
}

type EventstoreRepos struct {
	UserEvents    *usr_event.UserEventstore
	ProjectEvents *proj_event.ProjectEventstore
	OrgEvents     *org_events.OrgEventstore
	IamEvents     *iam_events.IAMEventstore
}

func Register(configs Configs, bulkLimit, errorCount uint64, view *view.View, eventstore eventstore.Eventstore, repos EventstoreRepos, systemDefaults sd.SystemDefaults) []spooler.Handler {
	return []spooler.Handler{
		&User{handler: handler{view, bulkLimit, configs.cycleDuration("User"), errorCount},
			orgEvents: repos.OrgEvents, iamEvents: repos.IamEvents, iamID: systemDefaults.IamID},
		&UserSession{handler: handler{view, bulkLimit, configs.cycleDuration("UserSession"), errorCount}, userEvents: repos.UserEvents},
		&UserMembership{handler: handler{view, bulkLimit, configs.cycleDuration("UserMembership"), errorCount}, orgEvents: repos.OrgEvents, projectEvents: repos.ProjectEvents},
		&Token{handler: handler{view, bulkLimit, configs.cycleDuration("Token"), errorCount}, ProjectEvents: repos.ProjectEvents},
		&Key{handler: handler{view, bulkLimit, configs.cycleDuration("Key"), errorCount}},
		&Application{handler: handler{view, bulkLimit, configs.cycleDuration("Application"), errorCount}, projectEvents: repos.ProjectEvents},
		&Org{handler: handler{view, bulkLimit, configs.cycleDuration("Org"), errorCount}},
		&UserGrant{
			handler:       handler{view, bulkLimit, configs.cycleDuration("UserGrant"), errorCount},
			eventstore:    eventstore,
			userEvents:    repos.UserEvents,
			orgEvents:     repos.OrgEvents,
			projectEvents: repos.ProjectEvents,
			iamEvents:     repos.IamEvents,
			iamID:         systemDefaults.IamID},
		&MachineKeys{handler: handler{view, bulkLimit, configs.cycleDuration("MachineKey"), errorCount}},
		&LoginPolicy{handler: handler{view, bulkLimit, configs.cycleDuration("LoginPolicy"), errorCount}},
		&IDPConfig{handler: handler{view, bulkLimit, configs.cycleDuration("IDPConfig"), errorCount}},
		&IDPProvider{handler: handler{view, bulkLimit, configs.cycleDuration("IDPProvider"), errorCount}, systemDefaults: systemDefaults, orgEvents: repos.OrgEvents, iamEvents: repos.IamEvents},
		&ExternalIDP{handler: handler{view, bulkLimit, configs.cycleDuration("ExternalIDP"), errorCount}, systemDefaults: systemDefaults, orgEvents: repos.OrgEvents, iamEvents: repos.IamEvents},
		&PasswordComplexityPolicy{handler: handler{view, bulkLimit, configs.cycleDuration("PasswordComplexityPolicy"), errorCount}},
		&OrgIAMPolicy{handler: handler{view, bulkLimit, configs.cycleDuration("OrgIAMPolicy"), errorCount}},
		&ProjectRole{handler: handler{view, bulkLimit, configs.cycleDuration("ProjectRole"), errorCount}, projectEvents: repos.ProjectEvents},
		spooler.NewGarbageCollector(view.Db, configs.cycleDuration("TokenCleanup"), tokenTable, user_view_model.TokenKeyExpiration),
		spooler.NewGarbageCollector(view.Db, configs.cycleDuration("KeyCleanup"), keyTable, key_view_model.KeyExpiry),
	}
}

func (configs Configs) cycleDuration(viewModel string) time.Duration {
	c, ok := configs[viewModel]
	if !ok {
		return 1 * time.Second
	}
	return c.MinimumCycleDuration.Duration
}

func (h *handler) MinimumCycleDuration() time.Duration {
	return h.cycleDuration
}

func (h *handler) QueryLimit() uint64 {
	return h.bulkLimit
}
