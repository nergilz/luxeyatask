package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

type IUsecase interface{}

type Handler struct {
	Logger   *slog.Logger
	ServeMux *http.ServeMux
	Service  IUsecase
}

func NewHandler(log *slog.Logger, service IUsecase) Handler {
	mux := http.NewServeMux()

	h := Handler{
		Logger:   log,
		ServeMux: mux,
	}

	h.addRoutes(mux)

	return h
}

func (h Handler) addRoutes(mux *http.ServeMux) {
	// для юзеров
	mux.Handle("/", h.handleHome())
	mux.Handle("/get_all", h.handleGetAll())
	mux.Handle("/registration", h.handleRegistration())
	mux.Handle("/set_score", h.handleSetScore())
	// турниры
	mux.Handle("/start_tour{name}", h.handleStartTour())
}

func (h Handler) handleHome() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Welcome to the Home Page!")

			h.Logger.Info("handle home page")
		},
	)
}

func (h Handler) handleGetAll() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "This is About Page.")

			h.Logger.Info("handle about page")
		},
	)
}

func (h Handler) handleRegistration() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "This is About Page.")

			h.Logger.Info("handle about page")
		},
	)
}

func (h Handler) handleSetScore() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "This is About Page.")

			h.Logger.Info("handle about page")
		},
	)
}

func (h Handler) handleStartTour() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "This is About Page.")

			h.Logger.Info("handle about page")
		},
	)
}
