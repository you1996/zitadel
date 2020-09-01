package handler

import (
	"net/http"

	"github.com/caos/zitadel/internal/auth_request/model"
)

const (
	tmplMfaInitDone = "MfaInitDone"
)

type mfaInitDoneData struct {
}

func (l *Login) renderMfaInitDone(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, data *mfaDoneData) {
	data.baseData = l.getBaseData(r, authReq, tmplMfaInitDone, nil)
	data.profileData = l.getProfileData(authReq)
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplMfaInitDone], data, nil)
}
