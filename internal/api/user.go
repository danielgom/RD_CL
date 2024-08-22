package api

import (
	"RD-Clone-NAPI/internal/dtos"
	services "RD-Clone-NAPI/internal/svc"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// UserHandler is an instance of our user handler API.
type UserHandler struct {
	api    *API
	UsrSvc services.UserService
}

// NewUserHandler returns a UserHandler instance.
func NewUserHandler(svc services.UserService, api *API) *UserHandler {
	return &UserHandler{UsrSvc: svc, api: api}
}

// SignUp is used to create a new user.
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req dtos.RegisterRequest

	err := shouldBindIntoAndValidate[dtos.RegisterRequest](r, &req, h.api.validator)
	if err != nil {
		renderAs(w, r, userError("invalid register request", err.Error()))
		return
	}

	response, err := h.UsrSvc.SignUp(r.Context(), &req)
	if err != nil {
		renderAs(w, r, internalServerError(err))
		return
	}

	renderJSON201(w, r, response)
}

// VerifyAccount verifies an account based on the token that has been given.
func (h *UserHandler) VerifyAccount(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	err := h.UsrSvc.VerifyAccount(r.Context(), token)
	if err != nil {
		renderAs(w, r, internalServerError(err))
		return
	}

	renderJSON200(w, r,
		map[string]any{"account": "Validated", "status": http.StatusOK})
}

// Login returns a JWT based on the user that has been logged in.
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dtos.LoginRequest

	err := shouldBindIntoAndValidate[dtos.LoginRequest](r, &req, h.api.validator)
	if err != nil {
		renderAs(w, r, userError("invalid login request", err.Error()))
		return
	}

	response, err := h.UsrSvc.Login(r.Context(), &req)
	if err != nil {
		renderAs(w, r, internalServerError(err))
		return
	}

	renderJSON201(w, r, response)
}

func (h *UserHandler) refreshToken(w http.ResponseWriter, r *http.Request) {
	var req dtos.RefreshTokenRequest

	err := shouldBindIntoAndValidate[dtos.RefreshTokenRequest](r, &req, h.api.validator)
	if err != nil {
		renderAs(w, r, userError("invalid refresh token request", err.Error()))
		return
	}

	response, err := h.UsrSvc.RefreshToken(r.Context(), &req)
	if err != nil {
		renderAs(w, r, internalServerError(err))
		return
	}

	renderJSON201(w, r, response)
}
