package iam

import (
	"time"

	"github.com/caos/zitadel/internal/v2/view"
)

type MemberSearchRequest struct {
	view.BaseSearchRequest
}

func (r *MemberSearchRequest) SortBy(key MemberSearchKey) *MemberSearchRequest {
	r.SortingColumn = &key
	return r
}

func (r *MemberSearchRequest) Filter(key MemberSearchKey, method view.SearchMethod, value interface{}) *MemberSearchRequest {
	if !key.Valid() {
		return r
	}
	r.BaseSearchRequest = *r.BaseSearchRequest.Filter(&key, method, value)
	return r
}

type MemberSearchKey int32

const (
	MemberSearchKeyUserName MemberSearchKey = iota + 1
	MemberSearchKeyEmail
	MemberSearchKeyFirstName
	MemberSearchKeyLastName
	MemberSearchKeyIamID
	MemberSearchKeyUserID

	memberSearchKeyCount
)

func (k *MemberSearchKey) ToColumnName() string {
	switch *k {
	case MemberSearchKeyUserName:
		return "user_name"
	case MemberSearchKeyEmail:
		return "email_address"
	case MemberSearchKeyFirstName:
		return "first_name"
	case MemberSearchKeyLastName:
		return "last_name"
	case MemberSearchKeyIamID:
		return "iam_id"
	case MemberSearchKeyUserID:
		return "user_id"
	default:
		return ""
	}
}

func (k *MemberSearchKey) Valid() bool {
	return *k > 0 && *k < memberSearchKeyCount
}

type MemberSearchResponse struct {
	Offset      uint64
	Limit       uint64
	TotalResult uint64
	Result      []*MemberView
	Sequence    uint64
	Timestamp   time.Time
}

func (r *MemberSearchRequest) EnsureLimit(limit uint64) {
	if r.Limit == 0 || r.Limit > limit {
		r.Limit = limit
	}
}
