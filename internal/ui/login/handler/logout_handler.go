package handler

import (
	"net/http"
)

const (
	tmplLogoutDone = "LogoutDone"
)

func (l *Login) handleLogoutDone(w http.ResponseWriter, r *http.Request) {
	l.renderLogoutDone(w, r)
}

func (l *Login) renderLogoutDone(w http.ResponseWriter, r *http.Request) {
	data := l.getUserData(r, nil, tmplLogoutDone, nil)
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplLogoutDone], data, nil)
}
