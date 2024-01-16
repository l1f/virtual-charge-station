package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/l1f/virtual-charge-station/internal/app"
	"github.com/l1f/virtual-charge-station/internal/data/models"
	"github.com/l1f/virtual-charge-station/internal/data/repositories"
	"github.com/l1f/virtual-charge-station/web/templates"
)

var stationId = 1

type dashboard struct {
	app               app.App
	stationRepository repositories.Station
}

func (d *dashboard) Dashboard(w http.ResponseWriter, r *http.Request) {
	stations, err := d.stationRepository.FetchAllStations()
	if err != nil {
		d.app.Logger.Error(err)
	}

	component := templates.Dashboard(*stations)

	component.Render(context.Background(), w)
}

func (d *dashboard) NewStation(w http.ResponseWriter, r *http.Request) {
	stationId = stationId + 1
	err := d.stationRepository.CreateStation(models.Station{
		UID:             fmt.Sprintf("STATION-%d", stationId),
		BackendWSUrl:    "wss://backend",
		BackendPassword: "BackendPassword",
	})

	if err != nil {
		d.app.Logger.Error(err)
	}
}
