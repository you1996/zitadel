import { Injectable } from '@angular/core';
import { Empty } from 'google-protobuf/google/protobuf/empty_pb';
import { Timestamp } from 'google-protobuf/google/protobuf/timestamp_pb';
import { BehaviorSubject } from 'rxjs';

import { IDPQuery } from '../proto/generated/zitadel/idp_pb';
import {
    AddMachineKeyRequest,
    AddMachineKeyResponse,
    AddOrgDomainRequest,
    AddOrgMemberRequest,
    ListHumanPasswordlessRequest,
    ListHumanPasswordlessResponse,
    ListLoginPolicyMultiFactorsResponse,
    ListOrgIDPsRequest,
    ListOrgIDPsResponse,
    RemoveHumanPasswordlessRequest,
    RemoveHumanPasswordlessResponse,
    RemoveOrgDomainRequest,
    RemoveOrgMemberRequest,
    UpdateMachineRequest,
    ListLoginPolicyMultiFactorsRequest,
    ValidateOrgDomainRequest,
    AddMultiFactorToLoginPolicyRequest,
    AddMultiFactorToLoginPolicyResponse,
    RemoveMultiFactorFromLoginPolicyRequest,
    RemoveMultiFactorFromLoginPolicyResponse,
    ListLoginPolicySecondFactorsResponse,
    AddSecondFactorToLoginPolicyResponse,
    AddSecondFactorToLoginPolicyRequest,
    UpdateOrgIDPResponse,
    AddOrgOIDCIDPRequest,
    RemoveSecondFactorFromLoginPolicyRequest,
    RemoveSecondFactorFromLoginPolicyResponse,
    GetLoginPolicyResponse,
    GetLoginPolicyRequest,
    UpdateCustomLoginPolicyRequest,
    UpdateCustomLoginPolicyResponse,
    GetOrgIDPByIDRequest,
    GetOrgIDPByIDResponse,
    AddCustomLoginPolicyRequest,
    AddCustomLoginPolicyResponse,
    ListMachineKeysRequest,
    ListMachineKeysResponse,
    ResetLoginPolicyToDefaultRequest,
    ResetLoginPolicyToDefaultResponse,
    AddIDPToLoginPolicyRequest,
    AddIDPToLoginPolicyResponse,
    RemoveIDPFromLoginPolicyRequest,
    ListLoginPolicyIDPsRequest,
    ListLoginPolicyIDPsResponse,
    UpdateOrgIDPRequest,
    AddOrgOIDCIDPResponse,
    UpdateOrgIDPOIDCConfigRequest,
    RemoveOrgIDPRequest,
    UpdateOrgIDPOIDCConfigResponse,
    RemoveOrgIDPResponse,
    ReactivateOrgIDPRequest,
    DeactivateOrgIDPRequest,
    AddHumanUserResponse,
    AddHumanUserRequest,
    AddMachineUserRequest,
    AddMachineUserResponse,
    UpdateMachineResponse,
    RemoveMachineKeyRequest,
    RemoveMachineKeyResponse,
    RemoveUserIDPRequest,
    RemoveUserIDPResponse,
    ListUserIDPsRequest,
    ListUserIDPsResponse,
    GetIAMResponse,
    GetIAMRequest,
    GetDefaultPasswordComplexityPolicyResponse,
    GetDefaultPasswordComplexityPolicyRequest,
    GetMyOrgRequest,
    GetMyOrgResponse,
    AddOrgDomainResponse,
    RemoveOrgDomainResponse,
    ListOrgDomainsRequest,
    ListOrgDomainsResponse,
    SetPrimaryOrgDomainRequest,
    SetPrimaryOrgDomainResponse,
    GenerateOrgDomainValidationResponse,
    GenerateOrgDomainValidationRequest,
    ValidateOrgDomainResponse,
    ListOrgMembersRequest,
    ListOrgMembersResponse,
    GetOrgByDomainGlobalResponse,
    GetOrgByDomainGlobalRequest,
    AddOrgResponse,
    AddOrgRequest,
    UpdateOrgMemberResponse,
    UpdateOrgMemberRequest,
    RemoveOrgMemberResponse,
    DeactivateOrgResponse,
    DeactivateOrgRequest,
    ReactivateOrgResponse,
    ReactivateOrgRequest,
    AddProjectGrantRequest,
    AddProjectGrantResponse,
    ListOrgMemberRolesResponse,
    ListOrgMemberRolesRequest,
    GetOrgIAMPolicyRequest,
    GetOrgIAMPolicyResponse,
    GetPasswordAgePolicyResponse,
    GetPasswordAgePolicyRequest,
    AddCustomPasswordAgePolicyRequest,
    AddCustomPasswordAgePolicyResponse,
    ResetPasswordAgePolicyToDefaultRequest,
    ResetPasswordAgePolicyToDefaultResponse,
    UpdateCustomPasswordAgePolicyRequest,
    UpdateCustomPasswordAgePolicyResponse,
    GetPasswordComplexityPolicyResponse,
    GetPasswordComplexityPolicyRequest,
    AddCustomPasswordComplexityPolicyRequest,
    AddCustomPasswordComplexityPolicyResponse,
    ResetPasswordComplexityPolicyToDefaultResponse,
    ResetPasswordComplexityPolicyToDefaultRequest,
    UpdateCustomPasswordComplexityPolicyResponse,
    UpdateCustomPasswordComplexityPolicyRequest,
    GetPasswordLockoutPolicyResponse,
    GetPasswordLockoutPolicyRequest,
    AddCustomPasswordLockoutPolicyRequest,
    AddCustomPasswordLockoutPolicyResponse,
    ResetPasswordLockoutPolicyToDefaultRequest,
    ResetPasswordLockoutPolicyToDefaultResponse,
    UpdateCustomPasswordLockoutPolicyResponse,
    UpdateCustomPasswordLockoutPolicyRequest,
    GetUserByIDRequest,
    GetUserByIDResponse,
    RemoveUserRequest,
    RemoveUserResponse,
    ListProjectMembersRequest,
    ListProjectMembersResponse,
    ListUserMembershipsRequest,
    ListUserMembershipsResponse,
    GetHumanProfileResponse,
    GetHumanProfileRequest,
    ListUserMultiFactorsResponse,
    ListUserMultiFactorsRequest,
    RemoveHumanMultiFactorOTPResponse,
    RemoveHumanMultiFactorOTPRequest,
    RemoveHumanMultiFactorU2FRequest,
    RemoveHumanMultiFactorU2FResponse,
    UpdateHumanProfileRequest,
    UpdateHumanProfileResponse,
    GetHumanEmailResponse,
    GetHumanEmailRequest,
    UpdateHumanEmailResponse,
    UpdateHumanEmailRequest,
    GetHumanPhoneResponse,
    GetHumanPhoneRequest,
    UpdateHumanPhoneResponse,
    UpdateHumanPhoneRequest,
    RemoveHumanPhoneRequest,
    DeactivateUserRequest,
    DeactivateUserResponse,
    AddUserGrantRequest,
    AddUserGrantResponse,
    ReactivateUserResponse,
    ReactivateUserRequest,
    AddProjectRoleRequest,
    GetHumanAddressResponse,
    GetHumanAddressRequest,
    ResendEmailVerificationRequest,
    ResendHumanInitializationRequest,
    ResendHumanInitializationResponse,
    ResendHumanPhoneVerificationRequest,
    ResendHumanPhoneVerificationResponse,
    SetHumanInitialPasswordRequest,
    SetHumanInitialPasswordResponse,
    SendHumanResetPasswordNotificationRequest,
    UpdateHumanAddressRequest,
    UpdateHumanAddressResponse,
    ListUsersRequest,
    ListUsersResponse,
    GetUserByLoginNameGlobalResponse,
    GetUserByLoginNameGlobalRequest,
    GetUserGrantByIDResponse,
    GetUserGrantByIDRequest,
    UpdateUserGrantRequest,
    UpdateUserGrantResponse,
    RemoveUserGrantRequest,
    RemoveUserGrantResponse,
    BulkRemoveUserGrantRequest,
    BulkRemoveUserGrantResponse,
    ListAppChangesRequest,
    ListOrgChangesRequest,
    ListOrgChangesResponse,
    ListProjectChangesResponse,
    ListProjectChangesRequest,
    ListUserChangesRequest,
    ListUserChangesResponse,
    ListProjectsResponse,
    ListProjectsRequest,
    ListGrantedProjectsRequest,
    ListGrantedProjectsResponse,
    GetOIDCInformationResponse,
    GetOIDCInformationRequest,
    GetProjectByIDRequest,
    GetProjectByIDResponse,
    GetGrantedProjectByIDRequest,
    GetGrantedProjectByIDResponse,
    AddProjectResponse,
    AddProjectRequest,
    UpdateProjectRequest,
    UpdateProjectResponse,
    UpdateProjectGrantRequest,
    UpdateProjectGrantResponse,
    RemoveProjectGrantResponse,
    RemoveProjectGrantRequest,
    DeactivateProjectResponse,
    DeactivateProjectRequest,
    ReactivateProjectResponse,
    ReactivateProjectRequest,
    ListProjectGrantsRequest,
    ListProjectGrantsResponse,
    ListProjectGrantMemberRolesResponse,
    ListProjectGrantMemberRolesRequest,
    AddProjectMemberRequest,
    AddProjectMemberResponse,
    UpdateProjectMemberRequest,
    UpdateProjectMemberResponse,
    AddProjectGrantMemberRequest,
    AddProjectGrantMemberResponse,
    UpdateProjectGrantMemberRequest,
    UpdateProjectGrantMemberResponse,
    ListProjectGrantMembersResponse,
    ListProjectGrantMembersRequest,
    RemoveAppRequest,
    RemoveAppResponse,
    UpdateOIDCAppConfigRequest,
    UpdateOIDCAppConfigResponse,
    UpdateAppRequest,
    UpdateAppResponse,
    AddOIDCAppRequest,
    AddOIDCAppResponse,
    ReactivateProjectGrantRequest,
    ReactivateProjectGrantResponse,
    DeactivateProjectGrantRequest,
    DeactivateProjectGrantResponse,
    RemoveProjectRequest,
    RemoveProjectResponse,
    GetProjectGrantByIDRequest,
    GetProjectGrantByIDResponse,
    ListProjectMemberRolesResponse,
    ListProjectMemberRolesRequest,
    GetAppByIDRequest,
    GetAppByIDResponse,
    ListAppsRequest,
    ListAppsResponse,
    RemoveProjectMemberRequest,
    RemoveProjectMemberResponse,
    UpdateProjectRoleResponse,
    UpdateProjectRoleRequest,
    RemoveProjectGrantMemberRequest,
    RemoveProjectGrantMemberResponse,
    ReactivateAppRequest,
    ReactivateAppResponse,
    DeactivateAppResponse,
    DeactivateAppRequest,
    RegenerateOIDCClientSecretRequest,
    RegenerateOIDCClientSecretResponse,
    ListProjectRolesRequest,
    ListProjectRolesResponse,
    AddProjectRoleResponse,
    BulkAddProjectRolesRequest,
    RemoveProjectRoleRequest,
    ListAppChangesResponse
} from '../proto/generated/zitadel/management_pb';
import { KeyType } from '../proto/generated/zitadel/auth_n_pb';
import { ListQuery } from '../proto/generated/zitadel/object_pb';
import { GrpcService } from './grpc.service';
import { DomainSearchQuery, DomainValidationType } from '../proto/generated/zitadel/org_pb';
import { PasswordComplexityPolicy, UserAddress } from '../proto/generated/management_pb';
import { Gender, MembershipQuery, SearchQuery } from '../proto/generated/zitadel/user_pb';
import { ListMyUserGrantsRequest, ListMyUserGrantsResponse } from '../proto/generated/zitadel/auth_pb';
import { ProjectQuery, RoleQuery } from '../proto/generated/zitadel/project_pb';
import { App, AppQuery } from '../proto/generated/zitadel/app_pb';

export type ResponseMapper<TResp, TMappedResp> = (resp: TResp) => TMappedResp;

@Injectable({
    providedIn: 'root',
})
export class ManagementService {
    public ownedProjectsCount: BehaviorSubject<number> = new BehaviorSubject(0);
    public grantedProjectsCount: BehaviorSubject<number> = new BehaviorSubject(0);

    constructor(private readonly grpcService: GrpcService) { }

    public listOrgIDPs(
        limit?: number,
        offset?: number,
        queryList?: IDPQuery[],
    ): Promise<ListOrgIDPsResponse> {
        const req = new ListOrgIDPsRequest();
        const metadata = new ListQuery();

        if (limit) {
            metadata.setLimit(limit);
        }
        if (offset) {
            metadata.setOffset(offset);
        }
        if (queryList) {
            req.setQueriesList(queryList);
        }
        return this.grpcService.mgmt.listOrgIDPs(req);
    }

    public listHumanPasswordless(userId: string): Promise<ListHumanPasswordlessResponse> {
        const req = new ListHumanPasswordlessRequest();
        req.setUserId(userId);
        return this.grpcService.mgmt.listHumanPasswordless(req);
    }

    public removeHumanPasswordless(tokenId: string, userId: string): Promise<RemoveHumanPasswordlessResponse> {
        const req = new RemoveHumanPasswordlessRequest();
        req.setTokenId(tokenId);
        req.setUserId(userId);
        return this.grpcService.mgmt.removeHumanPasswordless(req);
    }

    public listLoginPolicyMultiFactors(): Promise<ListLoginPolicyMultiFactorsResponse> {
        const req = new ListLoginPolicyMultiFactorsRequest();
        return this.grpcService.mgmt.listLoginPolicyMultiFactors(req);
    }

    public addMultiFactorToLoginPolicy(req: AddMultiFactorToLoginPolicyRequest): Promise<AddMultiFactorToLoginPolicyResponse> {
        return this.grpcService.mgmt.addMultiFactorToLoginPolicy(req);
    }

    public removeMultiFactorFromLoginPolicy(req: RemoveMultiFactorFromLoginPolicyRequest): Promise<RemoveMultiFactorFromLoginPolicyResponse> {
        return this.grpcService.mgmt.removeMultiFactorFromLoginPolicy(req);
    }

    public listLoginPolicySecondFactors(): Promise<ListLoginPolicySecondFactorsResponse> {
        const req = new Empty();
        return this.grpcService.mgmt.listLoginPolicySecondFactors(req);
    }

    public addSecondFactorToLoginPolicy(req: AddSecondFactorToLoginPolicyRequest): Promise<AddSecondFactorToLoginPolicyResponse> {
        return this.grpcService.mgmt.addSecondFactorToLoginPolicy(req);
    }

    public removeSecondFactorFromLoginPolicy(req: RemoveSecondFactorFromLoginPolicyRequest): Promise<RemoveSecondFactorFromLoginPolicyResponse> {
        return this.grpcService.mgmt.removeSecondFactorFromLoginPolicy(req);
    }

    public getLoginPolicy(): Promise<GetLoginPolicyResponse> {
        const req = new GetLoginPolicyRequest();
        return this.grpcService.mgmt.getLoginPolicy(req);
    }

    public updateCustomLoginPolicy(req: UpdateCustomLoginPolicyRequest): Promise<UpdateCustomLoginPolicyResponse> {
        return this.grpcService.mgmt.updateCustomLoginPolicy(req);
    }

    public addCustomLoginPolicy(req: AddCustomLoginPolicyRequest): Promise<AddCustomLoginPolicyResponse> {
        return this.grpcService.mgmt.addCustomLoginPolicy(req);
    }

    public resetLoginPolicyToDefault(): Promise<ResetLoginPolicyToDefaultResponse> {
        return this.grpcService.mgmt.resetLoginPolicyToDefault(new ResetLoginPolicyToDefaultRequest());
    }

    public addIDPToLoginPolicy(idpId: string): Promise<AddIDPToLoginPolicyResponse> {
        const req = new AddIDPToLoginPolicyRequest();
        req.setIdpId(idpId);
        return this.grpcService.mgmt.addIDPToLoginPolicy(req);
    }

    public removeIDPFromLoginPolicy(idpId: string): Promise<Empty> {
        const req = new RemoveIDPFromLoginPolicyRequest();
        req.setIdpId(idpId);
        return this.grpcService.mgmt.removeIDPFromLoginPolicy(req);
    }

    public listLoginPolicyIDPs(limit?: number, offset?: number): Promise<ListLoginPolicyIDPsResponse> {
        const req = new ListLoginPolicyIDPsRequest();
        const metadata = new ListQuery();
        if (limit) {
            metadata.setLimit(limit);
        }
        if (offset) {
            metadata.setOffset(offset);
        }
        return this.grpcService.mgmt.listLoginPolicyIDPs(req);
    }

    public getOrgIDPByID(
        id: string,
    ): Promise<GetOrgIDPByIDResponse> {
        const req = new GetOrgIDPByIDRequest();
        req.setId(id);
        return this.grpcService.mgmt.getOrgIDPByID(req);
    }

    public updateOrgIDP(
        req: UpdateOrgIDPRequest,
    ): Promise<UpdateOrgIDPResponse> {
        return this.grpcService.mgmt.updateOrgIDP(req);
    }

    public addOrgOIDCIDP(
        req: AddOrgOIDCIDPRequest,
    ): Promise<AddOrgOIDCIDPResponse> {
        return this.grpcService.mgmt.addOrgOIDCIDP(req);
    }

    public updateOrgIDPOIDCConfig(
        req: UpdateOrgIDPOIDCConfigRequest,
    ): Promise<UpdateOrgIDPOIDCConfigResponse> {
        return this.grpcService.mgmt.updateOrgIDPOIDCConfig(req);
    }

    public removeOrgIDP(
        idpId: string,
    ): Promise<RemoveOrgIDPResponse> {
        const req = new RemoveOrgIDPRequest();
        req.setIdpId(idpId);
        return this.grpcService.mgmt.removeOrgIDP(req);
    }

    public deactivateOrgIDP(
        idpId: string,
    ): Promise<Empty> {
        const req = new DeactivateOrgIDPRequest();
        req.setIdpId(idpId);
        return this.grpcService.mgmt.deactivateOrgIDP(req);
    }

    public reactivateOrgIDP(
        idpId: string,
    ): Promise<Empty> {
        const req = new ReactivateOrgIDPRequest();
        req.setIdpId(idpId);
        return this.grpcService.mgmt.reactivateOrgIDP(req);
    }

    public addHumanUser(request: AddHumanUserRequest): Promise<AddHumanUserResponse> {
        return this.grpcService.mgmt.addHumanUser(request);
    }

    public addMachineUser(request: AddMachineUserRequest): Promise<AddMachineUserResponse> {
        return this.grpcService.mgmt.addMachineUser(request);
    }

    public updateMachine(
        userId: string,
        name?: string,
        description?: string,
    ): Promise<UpdateMachineResponse> {
        const req = new UpdateMachineRequest();
        req.setUserId(userId);
        if (name) {
            req.setName(name);
        }
        if (description) {
            req.setDescription(description);
        }
        return this.grpcService.mgmt.updateMachine(req);
    }

    public addMachineKey(
        userId: string,
        type: KeyType,
        date?: Timestamp,
    ): Promise<AddMachineKeyResponse> {
        const req = new AddMachineKeyRequest();
        req.setType(type);
        req.setUserId(userId);
        if (date) {
            req.setExpirationDate(date);
        }
        return this.grpcService.mgmt.addMachineKey(req);
    }

    public removeMachineKey(
        keyId: string,
        userId: string,
    ): Promise<RemoveMachineKeyResponse> {
        const req = new RemoveMachineKeyRequest();
        req.setKeyId(keyId);
        req.setUserId(userId);

        return this.grpcService.mgmt.removeMachineKey(req);
    }

    public listMachineKeys(
        userId: string,
        limit?: number,
        offset?: number,
        asc?: boolean,
    ): Promise<ListMachineKeysResponse> {
        const req = new ListMachineKeysRequest();
        const metadata = new ListQuery();
        req.setUserId(userId);
        if (limit) {
            metadata.setLimit(limit);
        }
        if (offset) {
            metadata.setOffset(offset);
        }
        if (asc) {
            metadata.setAsc(asc);
        }
        req.setMetaData(metadata);
        return this.grpcService.mgmt.listMachineKeys(req);
    }

    public removeUserIDP(
        idpId: string,
        userId: string,
        linkedUserId: string,
    ): Promise<RemoveUserIDPResponse> {
        const req = new RemoveUserIDPRequest();
        req.setUserId(userId);
        req.setIdpId(idpId);
        req.setUserId(userId);
        req.setLinkedUserId(linkedUserId);
        return this.grpcService.mgmt.removeUserIDP(req);
    }

    public listUserIDPs(
        userId: string,
        limit?: number,
        offset?: number,
    ): Promise<ListUserIDPsResponse> {
        const req = new ListUserIDPsRequest();
        const metadata = new ListQuery();
        req.setUserId(userId);
        if (limit) {
            metadata.setLimit(limit);
        }
        if (offset) {
            metadata.setOffset(offset);
        }
        req.setMetaData(metadata);
        return this.grpcService.mgmt.listUserIDPs(req);
    }

    public getIAM(): Promise<GetIAMResponse> {
        const req = new GetIAMRequest();
        return this.grpcService.mgmt.getIAM(req);
    }

    public getDefaultPasswordComplexityPolicy(): Promise<GetDefaultPasswordComplexityPolicyResponse> {
        const req = new GetDefaultPasswordComplexityPolicyRequest();
        return this.grpcService.mgmt.getDefaultPasswordComplexityPolicy(req);
    }

    public getMyOrg(): Promise<GetMyOrgResponse> {
        const req = new GetMyOrgRequest();
        return this.grpcService.mgmt.getMyOrg(req);
    }

    public addOrgDomain(domain: string): Promise<AddOrgDomainResponse> {
        const req = new AddOrgDomainRequest();
        req.setDomain(domain);
        return this.grpcService.mgmt.addOrgDomain(req);
    }

    public removeOrgDomain(domain: string): Promise<RemoveOrgDomainResponse> {
        const req = new RemoveOrgDomainRequest();
        req.setDomain(domain);
        return this.grpcService.mgmt.removeOrgDomain(req);
    }

    public listOrgDomains(queryList?: DomainSearchQuery[]):
        Promise<ListOrgDomainsResponse> {
        const req: ListOrgDomainsRequest = new ListOrgDomainsRequest();
        // const metadata= new ListQuery();
        if (queryList) {
            req.setQueriesList(queryList);
        }
        return this.grpcService.mgmt.listOrgDomains(req);
    }

    public setPrimaryOrgDomain(domain: string): Promise<SetPrimaryOrgDomainResponse> {
        const req = new SetPrimaryOrgDomainRequest();
        req.setDomain(domain);
        return this.grpcService.mgmt.setPrimaryOrgDomain(req);
    }

    public generateOrgDomainValidation(domain: string, type: DomainValidationType):
        Promise<GenerateOrgDomainValidationResponse> {
        const req: GenerateOrgDomainValidationRequest = new GenerateOrgDomainValidationRequest();
        req.setDomain(domain);
        req.setType(type);

        return this.grpcService.mgmt.generateOrgDomainValidation(req);
    }

    public validateOrgDomain(domain: string):
        Promise<ValidateOrgDomainResponse> {
        const req = new ValidateOrgDomainRequest();
        req.setDomain(domain);

        return this.grpcService.mgmt.validateOrgDomain(req);
    }

    public listOrgMembers(limit: number, offset: number): Promise<ListOrgMembersResponse> {
        const req = new ListOrgMembersRequest();
        const query = new ListQuery();
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        req.setMetaData(query);

        return this.grpcService.mgmt.listOrgMembers(req);
    }

    public getOrgByDomainGlobal(domain: string): Promise<GetOrgByDomainGlobalResponse> {
        const req = new GetOrgByDomainGlobalRequest();
        req.setDomain(domain);
        return this.grpcService.mgmt.getOrgByDomainGlobal(req);
    }

    public addOrg(name: string): Promise<AddOrgResponse> {
        const req = new AddOrgRequest();
        req.setName(name);
        return this.grpcService.mgmt.addOrg(req);
    }

    public addOrgMember(userId: string, rolesList: string[]): Promise<Empty> {
        const req = new AddOrgMemberRequest();
        req.setUserId(userId);
        if (rolesList) {
            req.setRolesList(rolesList);
        }
        return this.grpcService.mgmt.addOrgMember(req);
    }

    public updateOrgMember(userId: string, rolesList: string[]): Promise<UpdateOrgMemberResponse> {
        const req = new UpdateOrgMemberRequest();
        req.setUserId(userId);
        req.setRolesList(rolesList);
        return this.grpcService.mgmt.updateOrgMember(req);
    }


    public removeOrgMember(userId: string): Promise<RemoveOrgMemberResponse> {
        const req = new RemoveOrgMemberRequest();
        req.setUserId(userId);
        return this.grpcService.mgmt.removeOrgMember(req);
    }

    public deactivateOrg(): Promise<DeactivateOrgResponse> {
        const req = new DeactivateOrgRequest();
        return this.grpcService.mgmt.deactivateOrg(req);
    }

    public reactivateOrg(): Promise<ReactivateOrgResponse> {
        const req = new ReactivateOrgRequest();
        return this.grpcService.mgmt.reactivateOrg(req);
    }

    public addProjectGrant(
        orgId: string,
        projectId: string,
        roleKeysList: string[],
    ): Promise<AddProjectGrantResponse> {
        const req = new AddProjectGrantRequest();
        req.setProjectId(projectId);
        req.setGrantedOrgId(orgId);
        req.setRoleKeysList(roleKeysList);
        return this.grpcService.mgmt.addProjectGrant(req);
    }

    public listOrgMemberRoles(): Promise<ListOrgMemberRolesResponse> {
        const req = new ListOrgMemberRolesRequest();
        return this.grpcService.mgmt.listOrgMemberRoles(req);
    }

    // Policy

    public getOrgIAMPolicy(): Promise<GetOrgIAMPolicyResponse> {
        const req = new GetOrgIAMPolicyRequest();
        return this.grpcService.mgmt.getOrgIAMPolicy(req);
    }

    public GetPasswordAgePolicy(): Promise<GetPasswordAgePolicyResponse> {
        const req = new GetPasswordAgePolicyRequest();
        return this.grpcService.mgmt.getPasswordAgePolicy(req);
    }

    public addCustomPasswordAgePolicy(
        maxAgeDays: number,
        expireWarnDays: number,
    ): Promise<AddCustomPasswordAgePolicyResponse> {
        const req = new AddCustomPasswordAgePolicyRequest();
        req.setMaxAgeDays(maxAgeDays);
        req.setExpireWarnDays(expireWarnDays);

        return this.grpcService.mgmt.addCustomPasswordAgePolicy(req);
    }

    public resetPasswordAgePolicyToDefault(): Promise<ResetPasswordAgePolicyToDefaultResponse> {
        const req = new ResetPasswordAgePolicyToDefaultRequest();
        return this.grpcService.mgmt.resetPasswordAgePolicyToDefault(req);
    }

    public updateCustomPasswordAgePolicy(
        maxAgeDays: number,
        expireWarnDays: number,
    ): Promise<UpdateCustomPasswordAgePolicyResponse> {
        const req = new UpdateCustomPasswordAgePolicyRequest();
        req.setMaxAgeDays(maxAgeDays);
        req.setExpireWarnDays(expireWarnDays);
        return this.grpcService.mgmt.updateCustomPasswordAgePolicy(req);
    }

    public GetPasswordComplexityPolicy(): Promise<GetPasswordComplexityPolicyResponse> {
        const req = new GetPasswordComplexityPolicyRequest();
        return this.grpcService.mgmt.getPasswordComplexityPolicy(req);
    }

    public addCustomPasswordComplexityPolicy(
        hasLowerCase: boolean,
        hasUpperCase: boolean,
        hasNumber: boolean,
        hasSymbol: boolean,
        minLength: number,
    ): Promise<AddCustomPasswordComplexityPolicyResponse> {
        const req = new AddCustomPasswordComplexityPolicyRequest();
        req.setHasLowercase(hasLowerCase);
        req.setHasUppercase(hasUpperCase);
        req.setHasNumber(hasNumber);
        req.setHasSymbol(hasSymbol);
        req.setMinLength(minLength);
        return this.grpcService.mgmt.addCustomPasswordComplexityPolicy(req);
    }

    public resetPasswordComplexityPolicyToDefault(): Promise<ResetPasswordComplexityPolicyToDefaultResponse> {
        const req = new ResetPasswordComplexityPolicyToDefaultRequest();
        return this.grpcService.mgmt.resetPasswordComplexityPolicyToDefault(req);
    }

    public updateCustomPasswordComplexityPolicy(
        hasLowerCase: boolean,
        hasUpperCase: boolean,
        hasNumber: boolean,
        hasSymbol: boolean,
        minLength: number,
    ): Promise<UpdateCustomPasswordComplexityPolicyResponse> {
        const req = new UpdateCustomPasswordComplexityPolicyRequest();
        req.setHasLowercase(hasLowerCase);
        req.setHasUppercase(hasUpperCase);
        req.setHasNumber(hasNumber);
        req.setHasSymbol(hasSymbol);
        req.setMinLength(minLength);
        return this.grpcService.mgmt.updateCustomPasswordComplexityPolicy(req);
    }

    public getPasswordLockoutPolicy(): Promise<GetPasswordLockoutPolicyResponse> {
        const req = new GetPasswordLockoutPolicyRequest();

        return this.grpcService.mgmt.getPasswordLockoutPolicy(req);
    }

    public addCustomPasswordLockoutPolicy(
        maxAttempts: number,
        showLockoutFailures: boolean,
    ): Promise<AddCustomPasswordLockoutPolicyResponse> {
        const req = new AddCustomPasswordLockoutPolicyRequest();
        req.setMaxAttempts(maxAttempts);
        req.setShowLockoutFailure(showLockoutFailures);

        return this.grpcService.mgmt.addCustomPasswordLockoutPolicy(req);
    }

    public resetPasswordLockoutPolicyToDefault(): Promise<ResetPasswordLockoutPolicyToDefaultResponse> {
        const req = new ResetPasswordLockoutPolicyToDefaultRequest();
        return this.grpcService.mgmt.resetPasswordLockoutPolicyToDefault(req);
    }

    public updateCustomPasswordLockoutPolicy(
        maxAttempts: number,
        showLockoutFailures: boolean,
    ): Promise<UpdateCustomPasswordLockoutPolicyResponse> {
        const req = new UpdateCustomPasswordLockoutPolicyRequest();
        req.setMaxAttempts(maxAttempts);
        req.setShowLockoutFailure(showLockoutFailures);
        return this.grpcService.mgmt.updateCustomPasswordLockoutPolicy(req);
    }

    public getLocalizedComplexityPolicyPatternErrorString(policy: PasswordComplexityPolicy.AsObject): string {
        if (policy.hasNumber && policy.hasSymbol) {
            return 'POLICY.PWD_COMPLEXITY.SYMBOLANDNUMBERERROR';
        } else if (policy.hasNumber) {
            return 'POLICY.PWD_COMPLEXITY.NUMBERERROR';
        } else if (policy.hasSymbol) {
            return 'POLICY.PWD_COMPLEXITY.SYMBOLERROR';
        } else {
            return 'POLICY.PWD_COMPLEXITY.PATTERNERROR';
        }
    }

    public getUserByID(id: string): Promise<GetUserByIDResponse> {
        const req = new GetUserByIDRequest();
        req.setId(id);
        return this.grpcService.mgmt.getUserByID(req);
    }

    public removeUser(id: string): Promise<RemoveUserResponse> {
        const req = new RemoveUserRequest();
        req.setId(id);
        return this.grpcService.mgmt.removeUser(req);
    }

    public listProjectMembers(
        projectId: string,
        limit: number,
        offset: number,
        queryList?: SearchQuery[],
    ): Promise<ListProjectMembersResponse> {
        const req = new ListProjectMembersRequest();
        const query = new ListQuery();
        req.setMetaData(query);
        req.setProjectId(projectId);
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        if (queryList) {
            req.setQueriesList(queryList);
        }
        req.setMetaData(query);
        return this.grpcService.mgmt.listProjectMembers(req);
    }

    public listUserMemberships(userId: string,
        limit: number, offset: number,
        queryList?: MembershipQuery[],
    ): Promise<ListUserMembershipsResponse> {
        const req = new ListUserMembershipsRequest();
        req.setUserId(userId);
        const metadata = new ListQuery();
        if (limit) {
            metadata.setLimit(limit);
        }
        if (offset) {
            metadata.setOffset(offset);
        }
        if (queryList) {
            req.setQueriesList(queryList);
        }
        req.setMetaData(metadata);
        return this.grpcService.mgmt.listUserMemberships(req);
    }

    public GetUserProfile(userId: string): Promise<GetHumanProfileResponse> {
        const req = new GetHumanProfileRequest();
        req.setUserId(userId);
        return this.grpcService.mgmt.getHumanProfile(req);
    }

    public listUserMultiFactors(userId: string): Promise<ListUserMultiFactorsResponse> {
        const req = new ListUserMultiFactorsRequest();
        req.setUserId(userId);
        return this.grpcService.mgmt.listUserMultiFactors(req);
    }

    public removeHumanMultiFactorOTP(userId: string): Promise<RemoveHumanMultiFactorOTPResponse> {
        const req = new RemoveHumanMultiFactorOTPRequest();
        req.setUserId(userId);
        return this.grpcService.mgmt.removeHumanMultiFactorOTP(req);
    }

    public removeHumanMultiFactorU2F(userId: string, id: string): Promise<RemoveHumanMultiFactorU2FResponse> {
        const req = new RemoveHumanMultiFactorU2FRequest();
        req.setUserId(userId);
        return this.grpcService.mgmt.removeHumanMultiFactorU2F(req);
    }

    public updateHumanProfile(
        userId: string,
        firstName?: string,
        lastName?: string,
        nickName?: string,
        preferredLanguage?: string,
        gender?: Gender,
    ): Promise<UpdateHumanProfileResponse> {
        const req = new UpdateHumanProfileRequest();
        req.setUserId(userId);
        if (firstName) {
            req.setFirstName(firstName);
        }
        if (lastName) {
            req.setLastName(lastName);
        }
        if (nickName) {
            req.setNickName(nickName);
        }
        if (gender) {
            req.setGender(gender);
        }
        if (preferredLanguage) {
            req.setPreferredLanguage(preferredLanguage);
        }
        return this.grpcService.mgmt.updateHumanProfile(req);
    }

    public getHumanEmail(id: string): Promise<GetHumanEmailResponse> {
        const req = new GetHumanEmailRequest();
        req.setUserId(id);
        return this.grpcService.mgmt.getHumanEmail(req);
    }

    public updateHumanEmail(userId: string, email: string): Promise<UpdateHumanEmailResponse> {
        const req = new UpdateHumanEmailRequest();
        req.setUserId(userId);
        req.setEmail(email);
        return this.grpcService.mgmt.updateHumanEmail(req);
    }

    public getHumanPhone(userId: string): Promise<GetHumanPhoneResponse> {
        const req = new GetHumanPhoneRequest();
        req.setUserId(userId);
        return this.grpcService.mgmt.getHumanPhone(req);
    }

    public updateHumanPhone(userId: string, phone: string): Promise<UpdateHumanPhoneResponse> {
        const req = new UpdateHumanPhoneRequest();
        req.setUserId(userId);
        req.setPhone(phone);
        return this.grpcService.mgmt.updateHumanPhone(req);
    }

    public removeHumanPhone(userId: string): Promise<Empty> {
        const req = new RemoveHumanPhoneRequest();
        req.setUserId(userId);
        return this.grpcService.mgmt.removeHumanPhone(req);
    }

    public deactivateUser(id: string): Promise<DeactivateUserResponse> {
        const req = new DeactivateUserRequest();
        req.setId(id);
        return this.grpcService.mgmt.deactivateUser(req);
    }

    public addUserGrant(
        userId: string,
        roleNamesList: string[],
        projectId?: string,
    ): Promise<AddUserGrantResponse> {
        const req = new AddUserGrantRequest();
        if (projectId) {
            req.setProjectId(projectId);
        }
        req.setUserId(userId);
        req.setRoleKeysList(roleNamesList);

        return this.grpcService.mgmt.addUserGrant(req);
    }

    public reactivateUser(id: string): Promise<ReactivateUserResponse> {
        const req = new ReactivateUserRequest();
        req.setId(id);
        return this.grpcService.mgmt.reactivateUser(req);
    }

    public AddRole(projectId: string, roleKey: string, displayName: string, group: string): Promise<AddProjectRoleResponse> {
        const req = new AddProjectRoleRequest();
        req.setProjectId(projectId);
        req.setRoleKey(roleKey);
        if (displayName) {
            req.setDisplayName(displayName);
        }
        req.setGroup(group);
        return this.grpcService.mgmt.addProjectRole(req);
    }

    public getHumanAddress(userId: string): Promise<GetHumanAddressResponse> {
        const req = new GetHumanAddressRequest();
        req.setUserId(userId);
        return this.grpcService.mgmt.getHumanAddress(req);
    }

    public resendHumanEmailVerification(userId: string): Promise<any> {
        const req = new ResendEmailVerificationRequest();
        req.setUserId(userId);
        return this.grpcService.mgmt.resendHumanEmailVerification(req);
    }

    public resendHumanInitialization(userId: string, newemail: string): Promise<ResendHumanInitializationResponse> {
        const req = new ResendHumanInitializationRequest();
        if (newemail) {
            req.setEmail(newemail);
        }
        req.setUserId(userId);

        return this.grpcService.mgmt.resendHumanInitialization(req);
    }

    public resendHumanPhoneVerification(userId: string): Promise<ResendHumanPhoneVerificationResponse> {
        const req = new ResendHumanPhoneVerificationRequest();
        req.setUserId(userId);
        return this.grpcService.mgmt.resendHumanPhoneVerification(req);
    }

    public setHumanInitialPassword(userId: string, password: string): Promise<SetHumanInitialPasswordResponse> {
        const req = new SetHumanInitialPasswordRequest();
        req.setUserId(userId);
        req.setPassword(password);
        return this.grpcService.mgmt.setHumanInitialPassword(req);
    }

    public sendHumanResetPasswordNotification(userId: string, type: SendHumanResetPasswordNotificationRequest.Type): Promise<any> {
        const req = new SendHumanResetPasswordNotificationRequest();
        req.setUserId(userId);
        req.setType(type);
        return this.grpcService.mgmt.sendHumanResetPasswordNotification(req);
    }

    public updateHumanAddress(address: UserAddress.AsObject): Promise<UpdateHumanAddressResponse> {
        const req = new UpdateHumanAddressRequest();
        req.setUserId(address.id);
        req.setStreetAddress(address.streetAddress);
        req.setPostalCode(address.postalCode);
        req.setLocality(address.locality);
        req.setRegion(address.region);
        req.setCountry(address.country);
        return this.grpcService.mgmt.updateHumanAddress(req);
    }

    public listUsers(limit: number, offset: number, queryList?: SearchQuery[]): Promise<ListUsersResponse> {
        const req = new ListUsersRequest();
        const query = new ListQuery();
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        req.setMetaData(query);
        if (queryList) {
            req.setQueriesList(queryList);
        }
        return this.grpcService.mgmt.listUsers(req);
    }

    public getUserByLoginNameGlobal(loginname: string): Promise<GetUserByLoginNameGlobalResponse> {
        const req = new GetUserByLoginNameGlobalRequest();
        req.setLoginName(loginname);
        return this.grpcService.mgmt.getUserByLoginNameGlobal(req);
    }

    // USER GRANTS

    public listUserGrants(
        limit?: number,
        offset?: number,
        queryList?: UserGrantSearchQuery[],
    ): Promise<ListMyUserGrantsResponse> {
        const req = new ListMyUserGrantsRequest();
        const query = new ListQuery();
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        if (queryList) {
            req.set(queryList);
        }
        req.setQuery(query);
        return this.grpcService.mgmt.listUserGrants(req);
    }


    public getUserGrantByID(
        grantId: string,
        userId: string,
    ): Promise<GetUserGrantByIDResponse> {
        const req = new GetUserGrantByIDRequest();
        req.setGrantId(grantId);
        req.setUserId(userId);

        return this.grpcService.mgmt.getUserGrantByID(req);
    }

    public UpdateUserGrant(
        grantId: string,
        userId: string,
        roleKeysList: string[],
    ): Promise<UpdateUserGrantResponse> {
        const req = new UpdateUserGrantRequest();
        req.setUserId(grantId);
        req.setRoleKeysList(roleKeysList);
        req.setUserId(userId);

        return this.grpcService.mgmt.updateUserGrant(req);
    }

    public RemoveUserGrant(
        grantId: string,
        userId: string,
    ): Promise<RemoveUserGrantResponse> {
        const req = new RemoveUserGrantRequest();
        req.setUserId(userId);
        req.setGrantId(grantId);

        return this.grpcService.mgmt.removeUserGrant(req);
    }

    public bulkRemoveUserGrant(
        grantIdsList: string[],
    ): Promise<BulkRemoveUserGrantResponse> {
        const req = new BulkRemoveUserGrantRequest();
        req.setGrantIdList(grantIdsList);

        return this.grpcService.mgmt.bulkRemoveUserGrant(req);
    }

    //

    public listAppChanges(appId: string, projectId: string, limit: number, offset: number): Promise<ListAppChangesResponse> {
        const req = new ListAppChangesRequest();
        const query = new ListQuery();
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        req.setAppId(appId);
        req.setProjectId(projectId);
        req.setQuery(query);
        return this.grpcService.mgmt.listAppChanges(req);
    }

    public listOrgChanges(id: string, limit: number, offset: number): Promise<ListOrgChangesResponse> {
        const req = new ListOrgChangesRequest();
        req.setOrgId(id);
        const query = new ListQuery();
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        req.setQuery(query);

        return this.grpcService.mgmt.listOrgChanges(req);
    }

    public ProjectChanges(projectId: string, limit: number, offset: number): Promise<ListProjectChangesResponse> {
        const req = new ListProjectChangesRequest();
        req.setProjectId(projectId);
        const query = new ListQuery();
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        req.setQuery(query);
        return this.grpcService.mgmt.listProjectChanges(req);
    }

    public listUserChanges(userId: string, limit: number, offset: number): Promise<ListUserChangesResponse> {
        const req = new ListUserChangesRequest();
        req.setUserId(userId);
        const query = new ListQuery();
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        req.setQuery(query);
        return this.grpcService.mgmt.listUserChanges(req);
    }

    // project

    public listProjects(
        limit?: number, offset?: number, queryList?: ProjectQuery[]): Promise<ListProjectsResponse> {
        const req = new ListProjectsRequest();
        const query = new ListQuery();
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        req.setMetaData(query);
        if (queryList) {
            req.setQueriesList(queryList);
        }
        return this.grpcService.mgmt.listProjects(req).then(value => {
            const count = value.toObject().resultList.length;
            if (count >= 0) {
                this.ownedProjectsCount.next(count);
            }

            return value;
        });
    }

    public listGrantedProjects(
        limit: number, offset: number, queryList?: ProjectQuery[]): Promise<ListGrantedProjectsResponse> {
        const req = new ListGrantedProjectsRequest();
        const query = new ListQuery();
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        req.setMetaData(query);
        if (queryList) {
            req.setQueriesList(queryList);
        }
        return this.grpcService.mgmt.listGrantedProjects(req).then(value => {
            this.grantedProjectsCount.next(value.toObject().resultList.length);
            return value;
        });
    }

    public getOIDCInformation(): Promise<GetOIDCInformationResponse> {
        const req = new GetOIDCInformationRequest();
        return this.grpcService.mgmt.getOIDCInformation(req);
    }

    public getProjectByID(projectId: string): Promise<GetProjectByIDResponse> {
        const req = new GetProjectByIDRequest();
        req.setId(projectId);
        return this.grpcService.mgmt.getProjectByID(req);
    }

    public getGrantedProjectByID(projectId: string, grantId: string): Promise<GetGrantedProjectByIDResponse> {
        const req = new GetGrantedProjectByIDRequest();
        req.setGrantId(grantId);
        req.setProjectId(projectId);
        return this.grpcService.mgmt.getGrantedProjectByID(req);
    }

    public addProject(project: AddProjectRequest.AsObject): Promise<AddProjectResponse> {
        const req = new AddProjectRequest();
        req.setName(project.name);
        return this.grpcService.mgmt.addProject(req).then(value => {
            const current = this.ownedProjectsCount.getValue();
            this.ownedProjectsCount.next(current + 1);
            return value;
        });
    }

    public updateProject(project: UpdateProjectRequest): Promise<UpdateProjectResponse> {
        return this.grpcService.mgmt.updateProject(project);
    }

    public updateProjectGrant(grantId: string, projectId: string, rolesList: string[]): Promise<UpdateProjectGrantResponse> {
        const req = new UpdateProjectGrantRequest();
        req.setRoleKeysList(rolesList);
        req.setGrantId(grantId);
        req.setProjectId(projectId);
        return this.grpcService.mgmt.updateProjectGrant(req);
    }

    public removeProjectGrant(grantId: string, projectId: string): Promise<RemoveProjectGrantResponse> {
        const req = new RemoveProjectGrantRequest();
        req.setGrantId(grantId);
        req.setProjectId(projectId);
        return this.grpcService.mgmt.removeProjectGrant(req);
    }

    public deactivateProject(projectId: string): Promise<DeactivateProjectResponse> {
        const req = new DeactivateProjectRequest();
        req.setId(projectId);
        return this.grpcService.mgmt.deactivateProject(req);
    }

    public ReactivateProject(projectId: string): Promise<ReactivateProjectResponse> {
        const req = new ReactivateProjectRequest();
        req.setId(projectId);
        return this.grpcService.mgmt.reactivateProject(req);
    }

    public listProjectGrants(projectId: string, limit: number, offset: number): Promise<ListProjectGrantsResponse> {
        const req = new ListProjectGrantsRequest();
        const query = new ListQuery();
        req.setProjectId(projectId);
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        req.setMetaData(query);
        return this.grpcService.mgmt.listProjectGrants(req);
    }

    public listProjectGrantMemberRoles(): Promise<ListProjectGrantMemberRolesResponse> {
        const req = new ListProjectGrantMemberRolesRequest();
        return this.grpcService.mgmt.listProjectGrantMemberRoles(req);
    }

    public addProjectMember(projectId: string, userId: string, rolesList: string[]): Promise<AddProjectMemberResponse> {
        const req = new AddProjectMemberRequest();
        req.setProjectId(projectId);
        req.setUserId(userId);
        req.setRolesList(rolesList);
        return this.grpcService.mgmt.addProjectMember(req);
    }

    public updateProjectMember(projectId: string, userId: string, rolesList: string[]): Promise<UpdateProjectMemberResponse> {
        const req = new UpdateProjectMemberRequest();
        req.setProjectId(projectId);
        req.setUserId(userId);
        req.setRolesList(rolesList);
        return this.grpcService.mgmt.updateProjectMember(req);
    }

    public addProjectGrantMember(
        projectId: string,
        grantId: string,
        userId: string,
        rolesList: string[],
    ): Promise<AddProjectGrantMemberResponse> {
        const req = new AddProjectGrantMemberRequest();
        req.setProjectId(projectId);
        req.setGrantId(grantId);
        req.setUserId(userId);
        req.setRolesList(rolesList);
        return this.grpcService.mgmt.addProjectGrantMember(req);
    }

    public updateProjectGrantMember(
        projectId: string,
        grantId: string,
        userId: string,
        rolesList: string[],
    ): Promise<UpdateProjectGrantMemberResponse> {
        const req = new UpdateProjectGrantMemberRequest();
        req.setProjectId(projectId);
        req.setGrantId(grantId);
        req.setUserId(userId);
        req.setRolesList(rolesList);
        return this.grpcService.mgmt.updateProjectGrantMember(req);
    }

    public listProjectGrantMembers(
        projectId: string,
        grantId: string,
        limit: number,
        offset: number,
        queryList?: SearchQuery[],
    ): Promise<ListProjectGrantMembersResponse> {
        const req = new ListProjectGrantMembersRequest();
        const query = new ListQuery();
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        if (queryList) {
            req.setQueriesList(queryList);
        }
        req.setProjectId(projectId);
        req.setGrantId(grantId);
        req.setMetaData(query);
        return this.grpcService.mgmt.listProjectGrantMembers(req);
    }

    public removeProjectGrantMember(
        projectId: string,
        grantId: string,
        userId: string,
    ): Promise<RemoveProjectGrantMemberResponse> {
        const req = new RemoveProjectGrantMemberRequest();
        req.setGrantId(grantId);
        req.setUserId(userId);
        req.setProjectId(projectId);
        return this.grpcService.mgmt.removeProjectGrantMember(req);
    }

    public reactivateApp(projectId: string, appId: string): Promise<ReactivateAppResponse> {
        const req = new ReactivateAppRequest();
        req.setAppId(appId);
        req.setProjectId(projectId);

        return this.grpcService.mgmt.reactivateApp(req);
    }

    public deactivateApp(projectId: string, appId: string): Promise<DeactivateAppResponse> {
        const req = new DeactivateAppRequest();
        req.setAppId(appId);
        req.setProjectId(projectId);

        return this.grpcService.mgmt.deactivateApp(req);
    }

    public RegenerateOIDCClientSecret(appId: string, projectId: string): Promise<RegenerateOIDCClientSecretResponse> {
        const req = new RegenerateOIDCClientSecretRequest();
        req.setAppId(appId);
        req.setProjectId(projectId);
        return this.grpcService.mgmt.regenerateOIDCClientSecret(req);
    }

    public listProjectRoles(
        projectId: string,
        limit: number,
        offset: number,
        queryList?: RoleQuery[],
    ): Promise<ListProjectRolesResponse> {
        const req = new ListProjectRolesRequest();
        req.setProjectId(projectId);
        const query = new ListQuery();
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        if (queryList) {
            req.setQueriesList(queryList);
        }
        req.setMetaData(query);
        return this.grpcService.mgmt.listProjectRoles(req);
    }

    public addProjectRole(req: AddProjectRoleRequest): Promise<AddProjectRoleResponse> {
        return this.grpcService.mgmt.addProjectRole(req);
    }

    public bulkAddProjectRoles(
        projectId: string,
        rolesList: BulkAddProjectRolesRequest.Role[],
    ): Promise<Empty> {
        const req = new BulkAddProjectRolesRequest();
        req.setProjectId(projectId);
        req.setRolesList(rolesList);
        return this.grpcService.mgmt.bulkAddProjectRoles(req);
    }

    public removeProjectRole(projectId: string, roleKey: string): Promise<Empty> {
        const req = new RemoveProjectRoleRequest();
        req.setProjectId(projectId);
        req.setRoleKey(roleKey);
        return this.grpcService.mgmt.removeProjectRole(req);
    }


    public updateProjectRole(projectId: string, roleKey: string, displayName: string, group: string):
        Promise<UpdateProjectRoleResponse> {
        const req = new UpdateProjectRoleRequest();
        req.setProjectId(projectId);
        req.setRoleKey(roleKey);
        req.setGroup(group);
        req.setDisplayName(displayName);
        return this.grpcService.mgmt.updateProjectRole(req);
    }


    public removeProjectMember(projectId: string, userId: string): Promise<RemoveProjectMemberResponse> {
        const req = new RemoveProjectMemberRequest();
        req.setProjectId(projectId);
        req.setUserId(userId);
        return this.grpcService.mgmt.removeProjectMember(req);
    }

    public listApps(
        projectId: string,
        limit: number,
        offset: number,
        queryList?: AppQuery[]): Promise<ListAppsResponse> {
        const req = new ListAppsRequest();
        const query = new ListQuery();
        req.setProjectId(projectId);
        if (limit) {
            query.setLimit(limit);
        }
        if (offset) {
            query.setOffset(offset);
        }
        if (queryList) {
            req.setQueriesList(queryList);
        }
        req.setMetaData(query);
        return this.grpcService.mgmt.listApps(req);
    }

    public getAppByID(projectId: string, applicationId: string): Promise<App.AsObject | undefined> {
        const req = new GetAppByIDRequest();
        req.setProjectId(projectId);
        req.setAppId(applicationId);
        const prom = this.grpcService.mgmt.getAppByID(req);
        return prom.then(app => {
            return app.toObject().app;
        });
    }

    public listProjectMemberRoles(): Promise<ListProjectMemberRolesResponse> {
        const req = new ListProjectMemberRolesRequest();
        return this.grpcService.mgmt.listProjectMemberRoles(req);
    }

    public getProjectGrantByID(grantId: string, projectId: string): Promise<GetProjectGrantByIDResponse> {
        const req = new GetProjectGrantByIDRequest();
        req.setGrantId(grantId);
        req.setProjectId(projectId);
        return this.grpcService.mgmt.getProjectGrantByID(req);
    }

    public RemoveProject(id: string): Promise<RemoveProjectResponse> {
        const req = new RemoveProjectRequest();
        req.setId(id);
        return this.grpcService.mgmt.removeProject(req).then(value => {
            const current = this.ownedProjectsCount.getValue();
            this.ownedProjectsCount.next(current > 0 ? current - 1 : 0);
            return value;
        });
    }


    public deactivateProjectGrant(grantId: string, projectId: string): Promise<DeactivateProjectGrantResponse> {
        const req = new DeactivateProjectGrantRequest();
        req.setGrantId(grantId);
        req.setProjectId(projectId);
        return this.grpcService.mgmt.deactivateProjectGrant(req);
    }

    public reactivateProjectGrant(grantId: string, projectId: string): Promise<ReactivateProjectGrantResponse> {
        const req = new ReactivateProjectGrantRequest();
        req.setGrantId(grantId);
        req.setProjectId(projectId);
        return this.grpcService.mgmt.reactivateProjectGrant(req);
    }

    public addOIDCApp(app: AddOIDCAppRequest): Promise<AddOIDCAppResponse> {
        return this.grpcService.mgmt.addOIDCApp(app);
    }

    public updateApp(projectId: string, appId: string, name: string): Promise<UpdateAppResponse> {
        const req = new UpdateAppRequest();
        req.setAppId(appId);
        req.setName(name);
        req.setProjectId(projectId);
        return this.grpcService.mgmt.updateApp(req);
    }

    public updateOIDCAppConfig(req: UpdateOIDCAppConfigRequest): Promise<UpdateOIDCAppConfigResponse> {
        return this.grpcService.mgmt.updateOIDCAppConfig(req);
    }

    public removeApp(projectId: string, appId: string): Promise<RemoveAppResponse> {
        const req = new RemoveAppRequest();
        req.setAppId(appId);
        req.setProjectId(projectId);
        return this.grpcService.mgmt.removeApp(req);
    }
}
