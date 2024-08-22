package api

import (
	"RD-Clone-NAPI/internal/config"
	services "RD-Clone-NAPI/internal/svc"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/render"
	"net/http"
)

func (a *API) Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(httplog.RequestLogger(routerLogger()))
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	factory := services.NewFactory()

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", getHealth)
		NewUserHandler(factory.UserService, a).Register(r)
	})

	return r
}

func (h *UserHandler) Register(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Post("/signup", h.SignUp)
		r.Get("/accountVerification/{token}", h.VerifyAccount)
		r.Post("/login", h.Login)
		r.Post("/refresh/token", h.refreshToken)
	})
}

func routerLogger() *httplog.Logger {
	logger := httplog.NewLogger("rd-clone-napi", httplog.Options{
		LogLevel:        config.Load().LogLevel(),
		JSON:            true,
		TimeFieldFormat: "2006-01-02 15:04:05.000000Z07:00",
		Tags: map[string]string{
			"env": config.Load().Server.Environment,
		},
	})

	return logger
}
