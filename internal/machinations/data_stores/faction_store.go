package datastores

import (
	repo "github.com/therobertcrocker/wayfinder/internal/machinations/repositories"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
)

type FactionStore struct {
	FactionRepo repo.FactionRepository
}

func NewFactionStore(storage db.StorageManager) *FactionStore {
	return &FactionStore{
		FactionRepo: repo.NewFactionRepo(storage),
	}
}
