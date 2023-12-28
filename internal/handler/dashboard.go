package handler

import (
	"context"
	"net/http"

	"github.com/l1f/virtual-charge-station/internal/data/models"
	"github.com/l1f/virtual-charge-station/web/templates"
)

var stations []models.Station = []models.Station{
	models.Station{
		Id:  0,
		UID: "TEST-0",
	},
	models.Station{
		Id:  1,
		UID: "TEST-1",
	},
	models.Station{
		Id:  2,
		UID: "TEST-2",
	},
	models.Station{
		Id:  3,
		UID: "TEST-3",
	},
	models.Station{
		Id:  4,
		UID: "TEST-4",
	},
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	component := templates.Dashboard(stations)

	component.Render(context.Background(), w)
}
