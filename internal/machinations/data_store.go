package machinations

import (
	repo "github.com/therobertcrocker/wayfinder/internal/machinations/repositories"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
)

// DataStore is a struct that contains all of the repositories for the Machinations domain
type DataStore struct {
	FactionRepo repo.FactionRepository
	AssetRepo   repo.AssetRepository
}

func NewDataStore(storage db.StorageManager) *DataStore {
	return &DataStore{
		FactionRepo: repo.NewFactionRepo(storage),
		AssetRepo:   repo.NewAssetRepo(storage),
	}
}
