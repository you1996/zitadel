package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	EndpointRoot             = "/"
	EndpointHealthz          = "/healthz"
	EndpointReadiness        = "/ready"
	EndpointLogin            = "/login"
	EndpointLoginName        = "/loginname"
	EndpointUserSelection    = "/userselection"
	EndpointChangeUsername   = "/username/change"
	EndpointPassword         = "/password"
	EndpointInitPassword     = "/password/init"
	EndpointChangePassword   = "/password/change"
	EndpointPasswordReset    = "/password/reset"
	EndpointInitUser         = "/user/init"
	EndpointMfaVerify        = "/mfa/verify"
	EndpointMfaPrompt        = "/mfa/prompt"
	EndpointMfaInitVerify    = "/mfa/init/verify"
	EndpointMailVerification = "/mail/verification"
	EndpointMailVerified     = "/mail/verified"
	EndpointRegister         = "/register"
	EndpointRegisterOrg      = "/register/org"
	EndpointLogoutDone       = "/logout/done"

	EndpointResources = "/resources"
)

func CreateRouter(login *Login, staticDir http.FileSystem, interceptors ...mux.MiddlewareFunc) *mux.Router {
	router := mux.NewRouter()
	router.Use(interceptors...)
	router.HandleFunc(EndpointRoot, login.handleLogin).Methods(http.MethodGet)
	router.HandleFunc(EndpointHealthz, login.handleHealthz).Methods(http.MethodGet)
	router.HandleFunc(EndpointReadiness, login.handleReadiness).Methods(http.MethodGet)
	router.HandleFunc(EndpointLogin, login.handleLogin).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc(EndpointLoginName, login.handleLoginName).Methods(http.MethodGet)
	router.HandleFunc(EndpointLoginName, login.handleLoginNameCheck).Methods(http.MethodPost)
	router.HandleFunc(EndpointUserSelection, login.handleSelectUser).Methods(http.MethodPost)
	router.HandleFunc(EndpointChangeUsername, login.handleChangeUsername).Methods(http.MethodPost)
	router.HandleFunc(EndpointPassword, login.handlePasswordCheck).Methods(http.MethodPost)
	router.HandleFunc(EndpointInitPassword, login.handlePasswordInit).Methods(http.MethodGet)
	router.HandleFunc(EndpointInitPassword, login.handlePasswordInitCheck).Methods(http.MethodPost)
	router.HandleFunc(EndpointPasswordReset, login.handlePasswordReset).Methods(http.MethodGet)
	router.HandleFunc(EndpointInitUser, login.handleUserInit).Methods(http.MethodGet)
	router.HandleFunc(EndpointInitUser, login.handleUserInitCheck).Methods(http.MethodPost)
	router.HandleFunc(EndpointMfaVerify, login.handleMfaVerification).Methods(http.MethodPost)
	router.HandleFunc(EndpointMfaPrompt, login.handleMfaPromptSelection).Methods(http.MethodGet)
	router.HandleFunc(EndpointMfaPrompt, login.handleMfaPrompt).Methods(http.MethodPost)
	router.HandleFunc(EndpointMfaInitVerify, login.handleMfaInitVerification).Methods(http.MethodPost)
	router.HandleFunc(EndpointMailVerification, login.handleMailVerification).Methods(http.MethodGet)
	router.HandleFunc(EndpointMailVerification, login.handleMailVerificationCheck).Methods(http.MethodPost)
	router.HandleFunc(EndpointChangePassword, login.handlePasswordChange).Methods(http.MethodPost)
	router.HandleFunc(EndpointRegister, login.handleRegistration).Methods(http.MethodGet)
	router.HandleFunc(EndpointRegister, login.handleRegistrationCheck).Methods(http.MethodPost)
	router.HandleFunc(EndpointLogoutDone, login.handleLogoutDone).Methods(http.MethodGet)
	router.PathPrefix(EndpointResources).Handler(login.handleResources(staticDir)).Methods(http.MethodGet)
	router.HandleFunc(EndpointRegisterOrg, login.handleRegistrationOrg).Methods(http.MethodGet)
	router.HandleFunc(EndpointRegisterOrg, login.handleRegistrationOrgCheck).Methods(http.MethodPost)
	return router
}
