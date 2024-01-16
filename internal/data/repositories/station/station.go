package station

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/google/uuid"
	"github.com/l1f/virtual-charge-station/internal/data"
	"github.com/l1f/virtual-charge-station/internal/data/models"
	"go.etcd.io/bbolt"
)

var bucketName = []byte("station")

type Station struct {
	DB *bbolt.DB
}

func (r *Station) FetchAllStations() (*[]models.Station, error) {
	stations := []models.Station{}

	err := r.DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf(data.ErrBucketNotFound, "station")
		}

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			station, err := decodeStation(v)
			if err != nil {
				return err
			}

			stations = append(stations, *station)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &stations, nil
}

func (r *Station) CreateStation(station models.Station) error {
	return r.DB.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf(data.ErrBucketNotFound, "station")
		}

		buffer, err := encodeStation(station)
		if err != nil {
			return err
		}

		uid, _ := uuid.NewRandom()
		err = bucket.Put([]byte(uid.String()), buffer.Bytes())

		return nil
	})
}

func decodeStation(encodedStation []byte) (*models.Station, error) {
	buffer := bytes.NewBuffer(encodedStation)
	decoder := gob.NewDecoder(buffer)
	station := models.Station{}

	if err := decoder.Decode(&station); err != nil {
		return nil, err
	}

	return &station, nil
}

func encodeStation(station models.Station) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)

	err := encoder.Encode(station)
	if err != nil {
		return nil, fmt.Errorf("Error while encode station data: %w", err)
	}

	return buffer, nil
}
