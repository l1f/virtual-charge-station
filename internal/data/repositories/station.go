package repositories

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

func encodeStation(encodedStation []byte) (*models.Station, error) {
	buffer := bytes.NewBuffer(encodedStation)
	decoder := gob.NewDecoder(buffer)
	station := models.Station{}

	if err := decoder.Decode(&station); err != nil {
		return nil, err
	}

	return &station, nil
}

func (r *Repositories) FetchAllStations() (*[]models.Station, error) {
	stations := []models.Station{}

	err := r.DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf(data.ErrBucketNotFound, "station")
		}

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			station, err := encodeStation(v)
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

	//TODO: if len of stations == 0, return error

	return &stations, nil
}

func (r *Repositories) CreateStation(station models.Station) error {
	return r.DB.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf(data.ErrBucketNotFound, "station")
		}

		buffer := new(bytes.Buffer)
		encoder := gob.NewEncoder(buffer)

		err := encoder.Encode(station)
		if err != nil {
			return fmt.Errorf("Error while encode station data: %w", err)
		}

		uid, _ := uuid.NewRandom()
		err = bucket.Put([]byte(uid.String()), buffer.Bytes())

		return nil
	})
}
