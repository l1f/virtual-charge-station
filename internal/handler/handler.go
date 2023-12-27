package handler

import "github.com/l1f/virtual-charge-station/internal/app"

type Handler struct {
	app *app.App
}

func New(app *app.App) *Handler {
	return &Handler{app: app}
}
