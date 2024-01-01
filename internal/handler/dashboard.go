package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/l1f/virtual-charge-station/internal/data/models"
	"github.com/l1f/virtual-charge-station/web/templates"
)

var stationId = 1

func (h *Handler) Dashboard(w http.ResponseWriter, r *http.Request) {
	stations, err := h.repositories.FetchAllStations()
	if err != nil {
		h.app.Logger.Error(err)
	}
	component := templates.Dashboard(*stations)

	component.Render(context.Background(), w)
}

func (h *Handler) NewStation(w http.ResponseWriter, r *http.Request) {
	stationId = stationId + 1
	err := h.repositories.CreateStation(models.Station{
		UID:             fmt.Sprintf("STATION-%d", stationId),
		BackendWSUrl:    "wss://backend",
		BackendPassword: "BackendPassword",
	})

	if err != nil {
		h.app.Logger.Error(err)
	}
}
