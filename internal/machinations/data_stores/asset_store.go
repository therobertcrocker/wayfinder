package datastores

import (
	"github.com/therobertcrocker/wayfinder/internal/machinations/domain"
	repo "github.com/therobertcrocker/wayfinder/internal/machinations/repositories"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
	"github.com/therobertcrocker/wayfinder/internal/util"
)

type AssetStore struct {
	AssetRepo repo.AssetRepository
	AssetPool *domain.AssetPool
}

func NewAssetStore(storage db.StorageManager) *AssetStore {
	return &AssetStore{
		AssetRepo: repo.NewAssetRepo(storage),
		AssetPool: &domain.AssetPool{
			Cunning: make(map[util.EntityID]*domain.Asset),
			Wealth:  make(map[util.EntityID]*domain.Asset),
			Force:   make(map[util.EntityID]*domain.Asset),
		},
	}
}

// LoadDataFromStorage loads all of the data from the storage layer into the data store
func (as *AssetStore) LoadDataFromStorage() error {
	if err := as.AssetRepo.LoadAssetsFromStorage(as.AssetPool); err != nil {
		return err
	}

	return nil
}

type AssetCSVRecord struct {
	Name      string `csv:"name"`
	Attribute string `csv:"attribute"`
}

// BulkAddAssets adds all of the assets to the storage layer
func (as *AssetStore) BulkAddAssets(filepath, fileType string) error {
	util.Log.Info("Adding assets to storage")

	return nil
}
