package query

import (
	"github.com/caos/zitadel/internal/eventstore/models"
)

type Handler interface {
	//spooler.Handler
	//ViewModel() string
	EventQuery() (*models.SearchQuery, error)
	Reduce(*models.Event) error
	OnError(event *models.Event, err error) error
	//OnSuccess() error
	//MinimumCycleDuration() time.Duration
	QueryLimit() uint64
}
