package handler

import (
	"net/http"

	"github.com/caos/zitadel/internal/auth_request/model"
)

const (
	tmplPasswordResetDone = "PasswordResetDone"
)

func (l *Login) handlePasswordReset(w http.ResponseWriter, r *http.Request) {
	authReq, err := l.getAuthRequest(r)
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	err = l.authRepo.RequestPasswordReset(setContext(r.Context(), authReq.UserOrgID), authReq.LoginName)
	l.renderPasswordResetDone(w, r, authReq, err)
}

func (l *Login) renderPasswordResetDone(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, err error) {
	data := l.getUserData(r, authReq, tmplPasswordResetDone, err)
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplPasswordResetDone], data, nil)
}
