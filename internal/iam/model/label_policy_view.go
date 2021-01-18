package model

import (
	"time"

	"github.com/caos/zitadel/internal/model"
)

type LabelPolicyView struct {
	AggregateID    string
	PrimaryColor   string
	SecondaryColor string
	WarnColor	   string
	Logo 		   []byte
	Border		   bool
	IconFont	   IconFont
	Font	       Font  
	Default        bool

	CreationDate time.Time
	ChangeDate   time.Time
	Sequence     uint64
}

type IconFont int32
const (
	IconFontDefault IconFont = iota
	IconFontLine
)

type Font int32
const (
	FontDefault Font = iota
	FontLATO
)


type LabelPolicySearchRequest struct {
	Offset        uint64
	Limit         uint64
	SortingColumn LabelPolicySearchKey
	Asc           bool
	Queries       []*LabelPolicySearchQuery
}

type LabelPolicySearchKey int32

const (
	LabelPolicySearchKeyUnspecified LabelPolicySearchKey = iota
	LabelPolicySearchKeyAggregateID
)

type LabelPolicySearchQuery struct {
	Key    LabelPolicySearchKey
	Method model.SearchMethod
	Value  interface{}
}

type LabelPolicySearchResponse struct {
	Offset      uint64
	Limit       uint64
	TotalResult uint64
	Result      []*LabelPolicyView
	Sequence    uint64
	Timestamp   time.Time
}
