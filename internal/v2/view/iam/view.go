package iam

import (
	"github.com/caos/zitadel/internal/eventstore/v2"
	"github.com/caos/zitadel/internal/v2/repository/iam"
	"github.com/lib/pq"
)

type MemberView struct {
	eventstore.View

	UserID string `gorm:"column:user_id"`
	IAMID  string `gorm:"column:iam_id"`
	// CreationDate time.Time      `gorm:"column:creation_date"`
	// ChangeDate   time.Time      `gorm:"column:change_date"`
	UserName     string         `gorm:"column:user_name"`
	EmailAddress string         `gorm:"column:email_address"`
	FirstName    string         `gorm:"column:first_name"`
	LastName     string         `gorm:"column:last_name"`
	Roles        pq.StringArray `gorm:"column:roles"`
	DisplayName  string         `gorm:"column:display_name"`
}

func (v *MemberView) Reduce() error {
	for _, event := range v.Events {
		switch e := event.(type) {
		case *iam.MemberAddedEvent:
			v.UserID = e.UserID
			v.IAMID = e.AggregateID()

		case *iam.MemberChangedEvent:
		case *iam.MemberRemovedEvent:
		}
		v.View.ReduceEvent(event)
	}

	return nil
}

func (v *MemberView) insert() error {
	return nil
}

func (v *MemberView) onError() error {
	return nil
}
