package model

import (
	"encoding/json"
	"time"

	org_es_model "github.com/caos/zitadel/internal/org/repository/eventsourcing/model"

	es_model "github.com/caos/zitadel/internal/iam/repository/eventsourcing/model"

	"github.com/caos/logging"
	caos_errs "github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore/models"
	"github.com/caos/zitadel/internal/iam/model"
)

const (
	LabelPolicyKeyAggregateID = "aggregate_id"
)

type LabelPolicyView struct {
	AggregateID  string    `json:"-" gorm:"column:aggregate_id;primary_key"`
	CreationDate time.Time `json:"-" gorm:"column:creation_date"`
	ChangeDate   time.Time `json:"-" gorm:"column:change_date"`
	State        int32     `json:"-" gorm:"column:label_policy_state"`

	PrimaryColor   string `json:"primaryColor" gorm:"column:primary_color"`
	SecondaryColor string `json:"secondaryColor" gorm:"column:secondary_color"`
	WarnColor	   string `json:"warnColor" gorm:"column:warn_color"`
	Border		   bool   `json:"border" gorm:"column:border"`
	Font           int32  `json:"font" gorm:"column:font"`
	IconFont       int32  `json:"iconFont" gorm:"column:iconFont"`
	Logo           []byte `json:"logo,omitempty" gorm:"column:logo"`
	Default        bool   `json:"-" gorm:"-"`

	Sequence uint64 `json:"-" gorm:"column:sequence"`
}

func LabelPolicyViewFromModel(policy *model.LabelPolicyView) *LabelPolicyView {
	return &LabelPolicyView{
		AggregateID:    policy.AggregateID,
		Sequence:       policy.Sequence,
		CreationDate:   policy.CreationDate,
		ChangeDate:     policy.ChangeDate,
		PrimaryColor:   policy.PrimaryColor,
		SecondaryColor: policy.SecondaryColor,
		WarnColor: 		policy.WarnColor,
		Logo:			policy.Logo,
		Border: 		policy.Border,
		Font: 			int32(policy.Font),
		IconFont: 		int32(policy.IconFont),
		Default:        policy.Default,
	}
}

func LabelPolicyViewToModel(policy *LabelPolicyView) *model.LabelPolicyView {
	return &model.LabelPolicyView{
		AggregateID:    policy.AggregateID,
		Sequence:       policy.Sequence,
		CreationDate:   policy.CreationDate,
		ChangeDate:     policy.ChangeDate,
		PrimaryColor:   policy.PrimaryColor,
		SecondaryColor: policy.SecondaryColor,
		WarnColor: 		policy.WarnColor,
		Border: 		policy.Border,
		Font: 			model.Font(policy.Font),
		IconFont: 		model.IconFont(policy.IconFont),
		Default:        policy.Default,
	}
}

func (i *LabelPolicyView) AppendEvent(event *models.Event) (err error) {
	i.Sequence = event.Sequence
	i.ChangeDate = event.CreationDate
	switch event.Type {
	case es_model.LabelPolicyAdded, org_es_model.LabelPolicyAdded:
		i.setRootData(event)
		i.CreationDate = event.CreationDate
		err = i.SetData(event)
	case es_model.LabelPolicyChanged, org_es_model.LabelPolicyChanged:
		err = i.SetData(event)
	}
	return err
}

func (r *LabelPolicyView) setRootData(event *models.Event) {
	r.AggregateID = event.AggregateID
}

func (r *LabelPolicyView) SetData(event *models.Event) error {
	if err := json.Unmarshal(event.Data, r); err != nil {
		logging.Log("MODEL-Flp9C").WithError(err).Error("could not unmarshal event data")
		return caos_errs.ThrowInternal(err, "MODEL-Hs8uf", "Could not unmarshal data")
	}
	return nil
}
