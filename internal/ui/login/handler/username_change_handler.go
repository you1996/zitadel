package handler

import (
	"net/http"

	"github.com/caos/zitadel/internal/auth_request/model"
)

const (
	tmplChangeUsername     = "ChangeUsername"
	tmplChangeUsernameDone = "ChangeUsernameDone"
)

type changeUsernameData struct {
	Username string `schema:"username"`
}

func (l *Login) renderChangeUsername(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, err error) {
	data := l.getUserData(r, authReq, tmplChangeUsername, err)
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplChangeUsername], data, nil)
}

func (l *Login) handleChangeUsername(w http.ResponseWriter, r *http.Request) {
	data := new(changeUsernameData)
	authReq, err := l.getAuthRequestAndParseData(r, data)
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	err = l.authRepo.ChangeUsername(setContext(r.Context(), authReq.UserOrgID), authReq.UserID, data.Username)
	if err != nil {
		l.renderChangeUsername(w, r, authReq, err)
		return
	}
	l.renderChangeUsernameDone(w, r, authReq)
}

func (l *Login) renderChangeUsernameDone(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest) {
	data := l.getUserData(r, authReq, tmplChangeUsernameDone, nil)
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplChangeUsernameDone], data, nil)
}
