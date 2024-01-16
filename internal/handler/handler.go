package handler

import (
	"github.com/l1f/virtual-charge-station/internal/app"
	"github.com/l1f/virtual-charge-station/internal/data/repositories/station"
)

type Handler struct {
	Dashboard dashboard
}

func New(app *app.App) *Handler {
	r := &station.Station{DB: app.Data.DB}

	return &Handler{
		Dashboard: dashboard{
			stationRepository: r,
		},
	}
}
