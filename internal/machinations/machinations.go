package machinations

import (
	db "github.com/therobertcrocker/wayfinder/internal/storage"
)

type Machinations struct {
	DataStore *DataStore
}

func NewMachinations(storage db.StorageManager) *Machinations {
	return &Machinations{
		DataStore: NewDataStore(storage),
	}
}
