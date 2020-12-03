package iam

import (
	"context"

	"github.com/caos/zitadel/internal/errors"
	caos_errs "github.com/caos/zitadel/internal/errors"
	iam_model "github.com/caos/zitadel/internal/iam/model"
	iam_view_model "github.com/caos/zitadel/internal/iam/repository/view/model"
	"github.com/caos/zitadel/internal/tracing"
	iam_repo "github.com/caos/zitadel/internal/v2/repository/iam"
	"github.com/caos/zitadel/internal/v2/view"
	"github.com/caos/zitadel/internal/v2/view/iam"
)

func (r *Repository) AddMember(ctx context.Context, member *iam_model.IAMMember) (*iam_model.IAMMember, error) {
	//TODO: check if roles valid

	if !member.IsValid() {
		return nil, caos_errs.ThrowPreconditionFailed(nil, "IAM-W8m4l", "Errors.IAM.MemberInvalid")
	}

	addedMember := iam_repo.NewMemberWriteModel(member.AggregateID, member.UserID)
	err := r.eventstore.FilterToQueryReducer(ctx, addedMember)
	if err != nil {
		return nil, err
	}
	if addedMember.Member.IsActive {
		return nil, errors.ThrowAlreadyExists(nil, "IAM-PtXi1", "Errors.IAM.Member.AlreadyExists")
	}

	iamAgg := iam_repo.AggregateFromWriteModel(&addedMember.WriteModel).
		PushMemberAdded(ctx, member.UserID, member.Roles...)

	err = r.eventstore.PushAggregate(ctx, addedMember, iamAgg)
	if err != nil {
		return nil, err
	}

	return writeModelToMember(addedMember), nil
}

//ChangeMember updates an existing member
func (r *Repository) ChangeMember(ctx context.Context, member *iam_model.IAMMember) (*iam_model.IAMMember, error) {
	//TODO: check if roles valid

	if !member.IsValid() {
		return nil, caos_errs.ThrowPreconditionFailed(nil, "IAM-LiaZi", "Errors.IAM.MemberInvalid")
	}

	existingMember, err := r.memberWriteModelByID(ctx, member.AggregateID, member.UserID)
	if err != nil {
		return nil, err
	}

	iam := iam_repo.AggregateFromWriteModel(&existingMember.WriteModel).
		PushMemberChangedFromExisting(ctx, existingMember, member.Roles...)

	err = r.eventstore.PushAggregate(ctx, existingMember, iam)
	if err != nil {
		return nil, err
	}

	return writeModelToMember(existingMember), nil
}

func (r *Repository) RemoveMember(ctx context.Context, member *iam_model.IAMMember) error {
	m, err := r.memberWriteModelByID(ctx, member.AggregateID, member.UserID)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	if errors.IsNotFound(err) {
		return nil
	}

	iamAgg := iam_repo.AggregateFromWriteModel(&m.WriteModel).
		PushEvents(iam_repo.NewMemberRemovedEvent(ctx, member.UserID))

	return r.eventstore.PushAggregate(ctx, m, iamAgg)
}

func (r *Repository) MemberByID(ctx context.Context, iamID, userID string) (member *iam_repo.MemberReadModel, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer func() { span.EndWithError(err) }()

	member = iam_repo.NewMemberReadModel(iamID, userID)
	err = r.eventstore.FilterToQueryReducer(ctx, member)
	if err != nil {
		return nil, err
	}

	return member, nil
}

func (r *Repository) SearchMember(ctx context.Context, search *iam_model.IAMMemberSearchRequest) (_ []*iam_view_model.IAMMemberView, count uint64, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer func() { span.EndWithError(err) }()

	members := []*iam.MemberView{}
	query := view.PrepareSearchQuery("adminapi.iam_members", memberSearchRequestFromIAMMemberSearchRequest(search))
	count, err = query(r.db, &members)
	if err != nil {
		return nil, 0, err
	}

	return memberViewsToIAMMemberViews(members), count, nil
}

func (r *Repository) memberWriteModelByID(ctx context.Context, iamID, userID string) (member *iam_repo.MemberWriteModel, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer func() { span.EndWithError(err) }()

	writeModel := iam_repo.NewMemberWriteModel(iamID, userID)
	err = r.eventstore.FilterToQueryReducer(ctx, writeModel)
	if err != nil {
		return nil, err
	}

	if !writeModel.Member.IsActive {
		return nil, errors.ThrowNotFound(nil, "IAM-D8JxR", "Errors.NotFound")
	}

	return writeModel, nil
}
