package handler

import (
	"net/http"
	"strconv"

	"github.com/caos/zitadel/internal/auth_request/model"
	caos_errs "github.com/caos/zitadel/internal/errors"
)

const (
	queryUserInitCode        = "code"
	queryUserInitUserID      = "userID"
	queryUserInitPasswordSet = "passwordset"

	tmplUserInit     = "UserInit"
	tmplUserInitDone = "UserInitDone"
)

type initUserFormData struct {
	Code            string `schema:"code"`
	Password        string `schema:"password"`
	PasswordConfirm string `schema:"passwordconfirm"`
	UserID          string `schema:"userID"`
	PasswordSet     bool   `schema:"passwordSet"`
	Resend          bool   `schema:"resend"`
}

type initUserData struct {
	baseData
	profileData
	Code                      string
	UserID                    string
	PasswordSet               bool
	PasswordPolicyDescription string
	MinLength                 uint64
	HasUppercase              string
	HasLowercase              string
	HasNumber                 string
	HasSymbol                 string
}

func (l *Login) handleUserInit(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue(queryUserInitUserID)
	code := r.FormValue(queryUserInitCode)
	passwordSet, _ := strconv.ParseBool(r.FormValue(queryUserInitPasswordSet))
	l.renderUserInit(w, r, nil, userID, code, passwordSet, nil)
}

func (l *Login) handleUserInitCheck(w http.ResponseWriter, r *http.Request) {
	data := new(initUserFormData)
	authReq, err := l.getAuthRequestAndParseData(r, data)
	if err != nil {
		l.renderError(w, r, nil, err)
		return
	}

	if data.Resend {
		l.resendUserInit(w, r, authReq, data.UserID, data.PasswordSet)
		return
	}
	l.checkUserInitCode(w, r, authReq, data, nil)
}

func (l *Login) checkUserInitCode(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, data *initUserFormData, err error) {
	if data.Password != data.PasswordConfirm {
		err := caos_errs.ThrowInvalidArgument(nil, "VIEW-fsdfd", "Errors.User.Password.ConfirmationWrong")
		l.renderUserInit(w, r, authReq, data.UserID, data.Code, data.PasswordSet, err)
		return
	}
	userOrgID := login
	if authReq != nil {
		userOrgID = authReq.UserOrgID
	}
	err = l.authRepo.VerifyInitCode(setContext(r.Context(), userOrgID), data.UserID, data.Code, data.Password)
	if err != nil {
		l.renderUserInit(w, r, authReq, data.UserID, "", data.PasswordSet, err)
		return
	}
	l.renderInitUserDone(w, r, authReq)
}

func (l *Login) resendUserInit(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, userID string, showPassword bool) {
	userOrgID := login
	if authReq != nil {
		userOrgID = authReq.UserOrgID
	}
	err := l.authRepo.ResendInitVerificationMail(setContext(r.Context(), userOrgID), userID)
	l.renderUserInit(w, r, authReq, userID, "", showPassword, err)
}

func (l *Login) renderUserInit(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, userID, code string, passwordSet bool, err error) {
	if authReq != nil {
		userID = authReq.UserID
	}
	data := initUserData{
		baseData:    l.getBaseData(r, authReq, tmplUserInit, err),
		profileData: l.getProfileData(authReq),
		UserID:      userID,
		Code:        code,
		PasswordSet: passwordSet,
	}
	policy, description, _ := l.getPasswordComplexityPolicyByUserID(r, userID)
	if policy != nil {
		data.PasswordPolicyDescription = description
		data.MinLength = policy.MinLength
		if policy.HasUppercase {
			data.HasUppercase = UpperCaseRegex
		}
		if policy.HasLowercase {
			data.HasLowercase = LowerCaseRegex
		}
		if policy.HasSymbol {
			data.HasSymbol = SymbolRegex
		}
		if policy.HasNumber {
			data.HasNumber = NumberRegex
		}
	}
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplUserInit], data, nil)
}

func (l *Login) renderInitUserDone(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest) {
	data := l.getUserData(r, authReq, tmplUserInitDone, nil)
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplUserInitDone], data, nil)
}
