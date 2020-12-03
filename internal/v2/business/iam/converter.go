package iam

import (
	es_models "github.com/caos/zitadel/internal/eventstore/models"
	"github.com/caos/zitadel/internal/eventstore/v2"
	old_iam_model "github.com/caos/zitadel/internal/iam/model"
	old_iam_view "github.com/caos/zitadel/internal/iam/repository/view/model"
	iam_repo "github.com/caos/zitadel/internal/v2/repository/iam"
	"github.com/caos/zitadel/internal/v2/repository/idp/oidc"
	"github.com/caos/zitadel/internal/v2/repository/member"
	"github.com/caos/zitadel/internal/v2/view"
	iam_view "github.com/caos/zitadel/internal/v2/view/iam"
)

func readModelToIAM(readModel *iam_repo.ReadModel) *old_iam_model.IAM {
	return &old_iam_model.IAM{
		ObjectRoot:                      readModelToObjectRoot(readModel.ReadModel),
		GlobalOrgID:                     readModel.GlobalOrgID,
		IAMProjectID:                    readModel.ProjectID,
		SetUpDone:                       old_iam_model.Step(readModel.SetUpDone),
		SetUpStarted:                    old_iam_model.Step(readModel.SetUpStarted),
		Members:                         readModelToMembers(&readModel.Members),
		DefaultLabelPolicy:              readModelToLabelPolicy(&readModel.DefaultLabelPolicy),
		DefaultLoginPolicy:              readModelToLoginPolicy(&readModel.DefaultLoginPolicy),
		DefaultOrgIAMPolicy:             readModelToOrgIAMPolicy(&readModel.DefaultOrgIAMPolicy),
		DefaultPasswordAgePolicy:        readModelToPasswordAgePolicy(&readModel.DefaultPasswordAgePolicy),
		DefaultPasswordComplexityPolicy: readModelToPasswordComplexityPolicy(&readModel.DefaultPasswordComplexityPolicy),
		DefaultPasswordLockoutPolicy:    readModelToPasswordLockoutPolicy(&readModel.DefaultPasswordLockoutPolicy),
		IDPs:                            readModelToIDPConfigs(&readModel.IDPs),
	}
}

func readModelToMembers(readModel *iam_repo.MembersReadModel) []*old_iam_model.IAMMember {
	members := make([]*old_iam_model.IAMMember, len(readModel.Members))

	for i, member := range readModel.Members {
		members[i] = &old_iam_model.IAMMember{
			ObjectRoot: readModelToObjectRoot(member.ReadModel),
			Roles:      member.Roles,
			UserID:     member.UserID,
		}
	}

	return members
}

func readModelToLabelPolicy(readModel *iam_repo.LabelPolicyReadModel) *old_iam_model.LabelPolicy {
	return &old_iam_model.LabelPolicy{
		ObjectRoot:     readModelToObjectRoot(readModel.ReadModel),
		PrimaryColor:   readModel.PrimaryColor,
		SecondaryColor: readModel.SecondaryColor,
		Default:        true,
		//TODO: State: int32,
	}
}

func readModelToLoginPolicy(readModel *iam_repo.LoginPolicyReadModel) *old_iam_model.LoginPolicy {
	return &old_iam_model.LoginPolicy{
		ObjectRoot:            readModelToObjectRoot(readModel.ReadModel),
		AllowExternalIdp:      readModel.AllowExternalIDP,
		AllowRegister:         readModel.AllowRegister,
		AllowUsernamePassword: readModel.AllowUserNamePassword,
		Default:               true,
		//TODO: IDPProviders: []*model.IDPProvider,
		//TODO: State: int32,
	}
}
func readModelToOrgIAMPolicy(readModel *iam_repo.OrgIAMPolicyReadModel) *old_iam_model.OrgIAMPolicy {
	return &old_iam_model.OrgIAMPolicy{
		ObjectRoot:            readModelToObjectRoot(readModel.ReadModel),
		UserLoginMustBeDomain: readModel.UserLoginMustBeDomain,
		Default:               true,
		//TODO: State: int32,
	}
}
func readModelToPasswordAgePolicy(readModel *iam_repo.PasswordAgePolicyReadModel) *old_iam_model.PasswordAgePolicy {
	return &old_iam_model.PasswordAgePolicy{
		ObjectRoot:     readModelToObjectRoot(readModel.ReadModel),
		ExpireWarnDays: uint64(readModel.ExpireWarnDays),
		MaxAgeDays:     uint64(readModel.MaxAgeDays),
		//TODO: State: int32,
	}
}
func readModelToPasswordComplexityPolicy(readModel *iam_repo.PasswordComplexityPolicyReadModel) *old_iam_model.PasswordComplexityPolicy {
	return &old_iam_model.PasswordComplexityPolicy{
		ObjectRoot:   readModelToObjectRoot(readModel.ReadModel),
		HasLowercase: readModel.HasLowercase,
		HasNumber:    readModel.HasNumber,
		HasSymbol:    readModel.HasSymbol,
		HasUppercase: readModel.HasUpperCase,
		MinLength:    uint64(readModel.MinLength),
		//TODO: State: int32,
	}
}
func readModelToPasswordLockoutPolicy(readModel *iam_repo.PasswordLockoutPolicyReadModel) *old_iam_model.PasswordLockoutPolicy {
	return &old_iam_model.PasswordLockoutPolicy{
		ObjectRoot:          readModelToObjectRoot(readModel.ReadModel),
		MaxAttempts:         uint64(readModel.MaxAttempts),
		ShowLockOutFailures: readModel.ShowLockOutFailures,
		//TODO: State: int32,
	}
}

func readModelToObjectRoot(readModel eventstore.ReadModel) es_models.ObjectRoot {
	return es_models.ObjectRoot{
		AggregateID:   readModel.AggregateID,
		ChangeDate:    readModel.ChangeDate,
		CreationDate:  readModel.CreationDate,
		ResourceOwner: readModel.ResourceOwner,
		Sequence:      readModel.ProcessedSequence,
	}
}

func writeModelToObjectRoot(writeModel eventstore.WriteModel) es_models.ObjectRoot {
	return es_models.ObjectRoot{
		AggregateID:   writeModel.AggregateID,
		ChangeDate:    writeModel.ChangeDate,
		ResourceOwner: writeModel.ResourceOwner,
		Sequence:      writeModel.ProcessedSequence,
	}
}

func readModelToMember(readModel *member.ReadModel) *old_iam_model.IAMMember {
	return &old_iam_model.IAMMember{
		ObjectRoot: readModelToObjectRoot(readModel.ReadModel),
		Roles:      readModel.Roles,
		UserID:     readModel.UserID,
	}
}

func writeModelToMember(writeModel *iam_repo.MemberWriteModel) *old_iam_model.IAMMember {
	return &old_iam_model.IAMMember{
		ObjectRoot: writeModelToObjectRoot(writeModel.Member.WriteModel),
		Roles:      writeModel.Member.Roles,
		UserID:     writeModel.Member.UserID,
	}
}

func readModelToIDPConfigView(rm *iam_repo.IDPConfigReadModel) *old_iam_model.IDPConfigView {
	return &old_iam_model.IDPConfigView{
		AggregateID:               rm.AggregateID,
		ChangeDate:                rm.ChangeDate,
		CreationDate:              rm.CreationDate,
		IDPConfigID:               rm.ConfigID,
		IDPProviderType:           old_iam_model.IDPProviderType(rm.ProviderType),
		IsOIDC:                    rm.OIDCConfig != nil,
		Name:                      rm.Name,
		OIDCClientID:              rm.OIDCConfig.ClientID,
		OIDCClientSecret:          rm.OIDCConfig.ClientSecret,
		OIDCIDPDisplayNameMapping: old_iam_model.OIDCMappingField(rm.OIDCConfig.IDPDisplayNameMapping),
		OIDCIssuer:                rm.OIDCConfig.Issuer,
		OIDCScopes:                rm.OIDCConfig.Scopes,
		OIDCUsernameMapping:       old_iam_model.OIDCMappingField(rm.OIDCConfig.UserNameMapping),
		Sequence:                  rm.ProcessedSequence,
		State:                     old_iam_model.IDPConfigState(rm.State),
		StylingType:               old_iam_model.IDPStylingType(rm.StylingType),
	}
}

func readModelToIDPConfigs(rm *iam_repo.IDPConfigsReadModel) []*old_iam_model.IDPConfig {
	configs := make([]*old_iam_model.IDPConfig, len(rm.Configs))
	for i, config := range rm.Configs {
		configs[i] = readModelToIDPConfig(&iam_repo.IDPConfigReadModel{ConfigReadModel: *config})
	}
	return configs
}

func readModelToIDPConfig(rm *iam_repo.IDPConfigReadModel) *old_iam_model.IDPConfig {
	return &old_iam_model.IDPConfig{
		ObjectRoot:  readModelToObjectRoot(rm.ReadModel),
		OIDCConfig:  readModelToIDPOIDCConfig(rm.OIDCConfig),
		IDPConfigID: rm.ConfigID,
		Name:        rm.Name,
		State:       old_iam_model.IDPConfigState(rm.State),
		StylingType: old_iam_model.IDPStylingType(rm.StylingType),
	}
}

func readModelToIDPOIDCConfig(rm *oidc.ConfigReadModel) *old_iam_model.OIDCIDPConfig {
	return &old_iam_model.OIDCIDPConfig{
		ObjectRoot:            readModelToObjectRoot(rm.ReadModel),
		ClientID:              rm.ClientID,
		ClientSecret:          rm.ClientSecret,
		ClientSecretString:    string(rm.ClientSecret.Crypted),
		IDPConfigID:           rm.IDPConfigID,
		IDPDisplayNameMapping: old_iam_model.OIDCMappingField(rm.IDPDisplayNameMapping),
		Issuer:                rm.Issuer,
		Scopes:                rm.Scopes,
		UsernameMapping:       old_iam_model.OIDCMappingField(rm.UserNameMapping),
	}
}

func writeModelToIDPConfig(wm *iam_repo.IDPConfigWriteModel) *old_iam_model.IDPConfig {
	return &old_iam_model.IDPConfig{
		ObjectRoot:  writeModelToObjectRoot(wm.WriteModel),
		OIDCConfig:  writeModelToIDPOIDCConfig(wm.OIDCConfig),
		IDPConfigID: wm.ConfigID,
		Name:        wm.Name,
		State:       old_iam_model.IDPConfigState(wm.State),
		StylingType: old_iam_model.IDPStylingType(wm.StylingType),
	}
}

func writeModelToIDPOIDCConfig(wm *oidc.ConfigWriteModel) *old_iam_model.OIDCIDPConfig {
	return &old_iam_model.OIDCIDPConfig{
		ObjectRoot:            writeModelToObjectRoot(wm.WriteModel),
		ClientID:              wm.ClientID,
		IDPConfigID:           wm.IDPConfigID,
		IDPDisplayNameMapping: old_iam_model.OIDCMappingField(wm.IDPDisplayNameMapping),
		Issuer:                wm.Issuer,
		Scopes:                wm.Scopes,
		UsernameMapping:       old_iam_model.OIDCMappingField(wm.UserNameMapping),
	}
}

func writeModelToIDPProvider(wm *iam_repo.LoginPolicyIDPProviderWriteModel) *old_iam_model.IDPProvider {
	return &old_iam_model.IDPProvider{
		ObjectRoot:  writeModelToObjectRoot(wm.WriteModel),
		IdpConfigID: wm.IDPConfigID,
		Type:        old_iam_model.IDPProviderType(wm.IDPProviderType),
	}
}

func memberSearchRequestFromIAMMemberSearchRequest(request *old_iam_model.IAMMemberSearchRequest) *iam_view.MemberSearchRequest {
	r := iam_view.NewMemberSearchRequest(request.Limit, request.Offset, request.Asc).
		SortBy(iam_view.MemberSearchKey(request.SortingColumn))
	for _, query := range request.Queries {
		r.Filter(iam_view.MemberSearchKey(query.Key), view.SearchMethod(query.Method), query.Value)
	}

	return r
}

func memberViewsToIAMMemberViews(m []*iam_view.MemberView) []*old_iam_view.IAMMemberView {
	members := make([]*old_iam_view.IAMMemberView, len(m))
	for i, member := range m {
		members[i] = memberViewToIAMMemberView(member)
	}
	return members
}

func memberViewToIAMMemberView(m *iam_view.MemberView) *old_iam_view.IAMMemberView {
	return &old_iam_view.IAMMemberView{
		UserID:       m.UserID,
		IAMID:        m.IAMID,
		UserName:     m.UserName,
		Email:        m.EmailAddress,
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		DisplayName:  m.DisplayName,
		Roles:        m.Roles,
		CreationDate: m.CreationDate,
		ChangeDate:   m.ChangeDate,
		Sequence:     m.ProcessedSequence,
	}
}
