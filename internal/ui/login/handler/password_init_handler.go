package handler

import (
	"net/http"

	"github.com/caos/zitadel/internal/auth_request/model"
	"github.com/caos/zitadel/internal/errors"
)

const (
	queryInitPWCode   = "code"
	queryInitPWUserID = "userID"

	tmplPasswordInit     = "PasswordInit"
	tmplPasswordInitDone = "PasswordInitDone"
)

type initPasswordFormData struct {
	Code            string `schema:"code"`
	Password        string `schema:"password"`
	PasswordConfirm string `schema:"passwordconfirm"`
	UserID          string `schema:"userID"`
	Resend          bool   `schema:"resend"`
}

type initPasswordData struct {
	baseData
	profileData
	Code                      string
	UserID                    string
	PasswordPolicyDescription string
	MinLength                 uint64
	HasUppercase              string
	HasLowercase              string
	HasNumber                 string
	HasSymbol                 string
}

func (l *Login) handlePasswordInit(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue(queryInitPWUserID)
	code := r.FormValue(queryInitPWCode)
	l.renderPasswordInit(w, r, nil, userID, code, nil)
}

func (l *Login) handlePasswordInitCheck(w http.ResponseWriter, r *http.Request) {
	data := new(initPasswordFormData)
	authReq, err := l.getAuthRequestAndParseData(r, data)
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}

	if data.Resend {
		l.resendPasswordSet(w, r, authReq)
		return
	}
	l.checkPWCode(w, r, authReq, data, nil)
}

func (l *Login) checkPWCode(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, data *initPasswordFormData, err error) {
	if data.Password != data.PasswordConfirm {
		err := errors.ThrowInvalidArgument(nil, "VIEW-KaGue", "Errors.User.Password.ConfirmationWrong")
		l.renderPasswordInit(w, r, authReq, data.UserID, data.Code, err)
		return
	}
	userOrg := login
	if authReq != nil {
		userOrg = authReq.UserOrgID
	}
	err = l.authRepo.SetPassword(setContext(r.Context(), userOrg), data.UserID, data.Code, data.Password)
	if err != nil {
		l.renderPasswordInit(w, r, authReq, data.UserID, "", err)
		return
	}
	l.renderPasswordInitDone(w, r, authReq)
}

func (l *Login) resendPasswordSet(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest) {
	userOrg := login
	if authReq != nil {
		userOrg = authReq.UserOrgID
	}
	err := l.authRepo.RequestPasswordReset(setContext(r.Context(), userOrg), authReq.LoginName)
	l.renderPasswordInit(w, r, authReq, authReq.UserID, "", err)
}

func (l *Login) renderPasswordInit(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, userID, code string, err error) {
	if userID == "" && authReq != nil {
		userID = authReq.UserID
	}

	data := initPasswordData{
		baseData:    l.getBaseData(r, authReq, tmplPasswordInit, err),
		profileData: l.getProfileData(authReq),
		UserID:      userID,
		Code:        code,
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
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplPasswordInit], data, nil)
}

func (l *Login) renderPasswordInitDone(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest) {
	data := l.getUserData(r, authReq, tmplPasswordInit, nil)
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplPasswordInitDone], data, nil)
}
