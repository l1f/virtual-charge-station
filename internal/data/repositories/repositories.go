package repositories

import (
	"go.etcd.io/bbolt"
)

type Repositories struct {
	DB *bbolt.DB
}

func New(db *bbolt.DB) *Repositories {
	return &Repositories{DB: db}
}
