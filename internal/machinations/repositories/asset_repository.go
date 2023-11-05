package repositories

import (
	"errors"

	"github.com/therobertcrocker/wayfinder/internal/machinations/entities"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
	"github.com/therobertcrocker/wayfinder/internal/util"
)

type AssetRepository interface {
	// Add Asset to the storage
	AddAsset(asset *entities.Asset) error

	// Add Assets to the storage
	AddAssets(assets []*entities.Asset) error

	// Get Asset from the storage
	GetAsset(id util.EntityID) (*entities.Asset, error)

	// Get Assets from the storage
	GetAssets(ids []util.EntityID) ([]*entities.Asset, error)

	// Update Asset in the storage
	UpdateAsset(asset *entities.Asset) error

	// Update Assets in the storage
	UpdateAssets(assets []*entities.Asset) error

	// Delete Asset from the storage
	DeleteAsset(id util.EntityID) error

	// Delete Assets from the storage
	DeleteAssets(ids []util.EntityID) error
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

// Add Asset to the storage
func (r *AssetRepo) AddAsset(asset *entities.Asset) error {
	//TODO: Implement
	return errors.New("not implemented")
}

// Add Assets to the storage
func (r *AssetRepo) AddAssets(assets []*entities.Asset) error {
	return errors.New("not implemented")
}

// Get Asset from the storage
func (r *AssetRepo) GetAsset(id util.EntityID) (*entities.Asset, error) {
	return nil, errors.New("not implemented")
}

// Get Assets from the storage
func (r *AssetRepo) GetAssets(ids []util.EntityID) ([]*entities.Asset, error) {
	return nil, errors.New("not implemented")
}

// Update Asset in the storage
func (r *AssetRepo) UpdateAsset(asset *entities.Asset) error {
	return errors.New("not implemented")
}

// Update Assets in the storage
func (r *AssetRepo) UpdateAssets(assets []*entities.Asset) error {
	return errors.New("not implemented")
}

// Delete Asset from the storage
func (r *AssetRepo) DeleteAsset(id util.EntityID) error {
	return errors.New("not implemented")
}

// Delete Assets from the storage
func (r *AssetRepo) DeleteAssets(ids []util.EntityID) error {
	return errors.New("not implemented")
}
