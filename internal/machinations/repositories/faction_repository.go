package repositories

import (
	db "github.com/therobertcrocker/wayfinder/internal/storage"
)

type FactionRepository interface {
}

var _ FactionRepository = (*FactionRepo)(nil)

type FactionRepo struct {
	Storage db.StorageManager
}

func NewFactionRepo(storage db.StorageManager) *FactionRepo {
	return &FactionRepo{
		Storage: storage,
	}
}
