package repositories

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/therobertcrocker/wayfinder/internal/common/logging"
	"github.com/therobertcrocker/wayfinder/internal/machinations/domain"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
)

const (
	AssetPoolBucket = "machinations_assetPool"
)

type AssetRepository interface {
	LoadAssetsFromStorage(pool *domain.AssetPool) error
	BulkAddAssets(pool *domain.AssetPool) error
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

func (ar *AssetRepo) BulkAddAssets(pool *domain.AssetPool) error {

	logging.Log.WithFields(log.Fields{"cunning": len(pool.Cunning), "wealth": len(pool.Wealth), "force": len(pool.Force)}).Debug("Adding assets to storage")

	cunningAssets, err := json.Marshal(pool.Cunning)
	if err != nil {
		return err
	}

	wealthAssets, err := json.Marshal(pool.Wealth)
	if err != nil {
		return err
	}

	forceAssets, err := json.Marshal(pool.Force)
	if err != nil {
		return err
	}

	if err := ar.Storage.Put(AssetPoolBucket, "cunning", cunningAssets); err != nil {
		return err
	}

	if err := ar.Storage.Put(AssetPoolBucket, "wealth", wealthAssets); err != nil {
		return err
	}

	if err := ar.Storage.Put(AssetPoolBucket, "force", forceAssets); err != nil {
		return err
	}

	logging.Log.Debug("Sucessfully added assets to storage")

	return nil
}
