package handler

import (
	"github.com/l1f/virtual-charge-station/internal/app"
	"github.com/l1f/virtual-charge-station/internal/data/repositories"
)

type Handler struct {
	app          *app.App
	repositories repositories.Repositories
}

func New(app *app.App) *Handler {
	r := repositories.New(app.Data.DB)

	return &Handler{
		app:          app,
		repositories: *r,
	}
}
