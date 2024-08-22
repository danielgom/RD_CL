package api

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type Error struct {
	Err  error `json:"-"`
	Code int   `json:"-"`

	RequestID string `json:"request_id,omitempty"`
	Message   string `json:"error,omitempty"`
	Reason    string `json:"reason,omitempty"`
}

func (e *Error) Render(_ http.ResponseWriter, r *http.Request) error {
	e.RequestID = middleware.GetReqID(r.Context())

	render.Status(r, e.Code)

	return nil
}

func internalServerError(err error) *Error {
	slog.Error("request failed", "error", err)
	return &Error{
		Err:     err,
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
}

func userError(msg, reason string) *Error {
	slog.Warn("invalid request", "msg", msg)
	return &Error{
		Code:    http.StatusBadRequest,
		Message: msg,
		Reason:  reason,
	}
}
