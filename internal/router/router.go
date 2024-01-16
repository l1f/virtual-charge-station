package router

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/l1f/virtual-charge-station/internal/app"
	"github.com/l1f/virtual-charge-station/internal/handler"
)

func getRouter(app *app.App) *chi.Mux {
	var router *chi.Mux = chi.NewRouter()
	var handlers *handler.Handler = handler.New(app)

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logger(app.Logger, &Options{
		WithUserAgent: true,
		WithReferer:   true,
	}))
	router.Use(middleware.Recoverer)

	// TODO: build in to binary
	router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))

	router.NotFound(handlers.NotFound)
	router.Get("/dashboard", handlers.Dashboard.Dashboard)

	router.Get("/newStation", handlers.Dashboard.NewStation)

	return router
}

func Start(app *app.App) error {
	router := getRouter(app)

	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	return server.ListenAndServe()
}
