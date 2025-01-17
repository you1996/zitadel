package model

import (
	"encoding/json"
	"time"

	"github.com/caos/logging"
	"github.com/lib/pq"

	"github.com/caos/zitadel/internal/domain"
	caos_errs "github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore/v1/models"
	"github.com/caos/zitadel/internal/iam/model"
	es_model "github.com/caos/zitadel/internal/iam/repository/eventsourcing/model"
)

const (
	IAMMemberKeyUserID    = "user_id"
	IAMMemberKeyIamID     = "iam_id"
	IAMMemberKeyUserName  = "user_name"
	IAMMemberKeyEmail     = "email"
	IAMMemberKeyFirstName = "first_name"
	IAMMemberKeyLastName  = "last_name"
)

type IAMMemberView struct {
	UserID             string         `json:"userId" gorm:"column:user_id;primary_key"`
	IAMID              string         `json:"-" gorm:"column:iam_id"`
	UserName           string         `json:"-" gorm:"column:user_name"`
	Email              string         `json:"-" gorm:"column:email_address"`
	FirstName          string         `json:"-" gorm:"column:first_name"`
	LastName           string         `json:"-" gorm:"column:last_name"`
	DisplayName        string         `json:"-" gorm:"column:display_name"`
	Roles              pq.StringArray `json:"roles" gorm:"column:roles"`
	Sequence           uint64         `json:"-" gorm:"column:sequence"`
	PreferredLoginName string         `json:"-" gorm:"column:preferred_login_name"`
	AvatarKey          string         `json:"-" gorm:"column:avatar_key"`
	UserResourceOwner  string         `json:"-" gorm:"column:user_resource_owner"`

	CreationDate time.Time `json:"-" gorm:"column:creation_date"`
	ChangeDate   time.Time `json:"-" gorm:"column:change_date"`
}

func IAMMemberToModel(member *IAMMemberView, prefixAvatarURL string) *model.IAMMemberView {
	return &model.IAMMemberView{
		UserID:             member.UserID,
		IAMID:              member.IAMID,
		UserName:           member.UserName,
		Email:              member.Email,
		FirstName:          member.FirstName,
		LastName:           member.LastName,
		DisplayName:        member.DisplayName,
		PreferredLoginName: member.PreferredLoginName,
		AvatarURL:          domain.AvatarURL(prefixAvatarURL, member.UserResourceOwner, member.AvatarKey),
		UserResourceOwner:  member.UserResourceOwner,
		Roles:              member.Roles,
		Sequence:           member.Sequence,
		CreationDate:       member.CreationDate,
		ChangeDate:         member.ChangeDate,
	}
}

func IAMMembersToModel(roles []*IAMMemberView, prefixAvatarURL string) []*model.IAMMemberView {
	result := make([]*model.IAMMemberView, len(roles))
	for i, r := range roles {
		result[i] = IAMMemberToModel(r, prefixAvatarURL)
	}
	return result
}

func (r *IAMMemberView) AppendEvent(event *models.Event) (err error) {
	r.Sequence = event.Sequence
	r.ChangeDate = event.CreationDate
	switch event.Type {
	case es_model.IAMMemberAdded:
		r.setRootData(event)
		r.CreationDate = event.CreationDate
		err = r.SetData(event)
	case es_model.IAMMemberChanged:
		err = r.SetData(event)
	}
	return err
}

func (r *IAMMemberView) setRootData(event *models.Event) {
	r.IAMID = event.AggregateID
}

func (r *IAMMemberView) SetData(event *models.Event) error {
	if err := json.Unmarshal(event.Data, r); err != nil {
		logging.Log("EVEN-Psl89").WithError(err).Error("could not unmarshal event data")
		return caos_errs.ThrowInternal(err, "MODEL-lub6s", "Could not unmarshal data")
	}
	return nil
}
