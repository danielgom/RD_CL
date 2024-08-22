package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

const (
	status200 = 200
	status201 = 201
)

func renderAs(w http.ResponseWriter, r *http.Request, v render.Renderer) {
	err := render.Render(w, r, v)
	if err != nil {
		slog.Error("Failed to render", "body", v, "error", err)
		err = render.Render(w, r, internalServerError(err))
		if err != nil {
			panic("Failed to render internal server error: " + err.Error())
		}
	}
}

func renderJSON200(w http.ResponseWriter, r *http.Request, obj any) {
	reqInfo := fmt.Sprintf("Path: %s ", r.URL.Path)
	slog.Debug(reqInfo, "body", obj)
	render.Status(r, status200)
	render.JSON(w, r, obj)
}

func renderJSON201(w http.ResponseWriter, r *http.Request, obj any) {
	reqInfo := fmt.Sprintf("Path: %s ", r.URL.Path)
	slog.Debug(reqInfo, "body", obj)
	render.Status(r, status201)
	render.JSON(w, r, obj)
}
