package spooler

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	caos_errs "github.com/caos/zitadel/internal/errors"
)

type GarbageCollector interface {
	Handler
	CleanUp() error
}

type garbageCollector struct {
	db               *gorm.DB
	cycleDuration    time.Duration
	viewName         string
	expirationColumn string
}

func (g *garbageCollector) ViewModel() string {
	return g.viewName
}

func (g *garbageCollector) OnSuccess() error {
	return nil
}

func (g *garbageCollector) MinimumCycleDuration() time.Duration {
	return g.cycleDuration
}

func (g *garbageCollector) CleanUp() error {
	err := g.db.Table(g.viewName).
		Where(fmt.Sprintf("%s < ?", g.expirationColumn), time.Now().UTC()).
		Delete(nil).
		Error
	if err != nil {
		return caos_errs.ThrowInternal(err, "SPOOL-ad2k1", "could not delete expired objects")
	}
	return nil
}

func NewGarbageCollector(db *gorm.DB, cycleDuration time.Duration, viewName, expirationColumn string) GarbageCollector {
	return &garbageCollector{
		db:               db,
		viewName:         viewName,
		cycleDuration:    cycleDuration,
		expirationColumn: expirationColumn,
	}
}
