package repositories

import (
	"encoding/json"

	"github.com/therobertcrocker/wayfinder/internal/machinations/domain"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
)

const (
	AssetPoolBucket = "machinations_assetPool"
)

type AssetRepository interface {
	LoadAssetsFromStorage(pool *domain.AssetPool) error
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

func (ar *AssetRepo) LoadAssetsFromStorage(pool *domain.AssetPool) error {
	cunningAssets, err := ar.Storage.Get(AssetPoolBucket, "cunning")
	if err != nil {
		return err
	}

	wealthAssets, err := ar.Storage.Get(AssetPoolBucket, "wealth")
	if err != nil {
		return err
	}

	forceAssets, err := ar.Storage.Get(AssetPoolBucket, "force")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(cunningAssets, &pool.Cunning); err != nil {
		return err
	}

	if err := json.Unmarshal(wealthAssets, &pool.Wealth); err != nil {
		return err
	}

	if err := json.Unmarshal(forceAssets, &pool.Force); err != nil {
		return err
	}

	return nil
}
