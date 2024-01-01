package data

import (
	"fmt"
	"go.etcd.io/bbolt"
)

var (
	ErrBucketNotFound = "Bucket %s not found"
)

// maybe not the best solution(?)
var buckets []string = []string{
	"station",
}

type Data struct {
	DB *bbolt.DB
}

func Open(path string) (*Data, error) {
	db, err := bbolt.Open(path, 0600, bbolt.DefaultOptions)
	if err != nil {
		return nil, err
	}

	return &Data{DB: db}, nil
}

func (r *Data) PrepareDB() error {
	err := r.DB.Update(func(tx *bbolt.Tx) error {
		for _, bucket := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(bucket))
			if err != nil {
				return fmt.Errorf("could not create %s bucket: %v", bucket, err)
			}
		}

		return nil
	})

	return err
}
