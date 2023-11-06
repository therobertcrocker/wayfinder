package repositories

import (
	"errors"

	"github.com/therobertcrocker/wayfinder/internal/machinations/entities"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
)

type AssetRepository interface {

	// Add Assets to the storage
	AddAssets(assets []*entities.Asset) error
}

var _ AssetRepository = (*AssetRepo)(nil)

type AssetRepo struct {
	Storage db.StorageManager
}

func NewAssetRepo(storage db.StorageManager) *AssetRepo {
	return &AssetRepo{
		Storage: storage,
	}
}

// Add Assets to the storage
func (r *AssetRepo) AddAssets(assets []*entities.Asset) error {
	return errors.New("not implemented")
}
