package repositories

import "github.com/l1f/virtual-charge-station/internal/data/models"

type Station interface {
	FetchAllStations() (*[]models.Station, error)
	CreateStation(station models.Station) error
}
