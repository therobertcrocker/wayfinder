package datastores

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/therobertcrocker/wayfinder/internal/common/logging"
	"github.com/therobertcrocker/wayfinder/internal/common/util"
	"github.com/therobertcrocker/wayfinder/internal/machinations/data_stores/loader"
	"github.com/therobertcrocker/wayfinder/internal/machinations/domain"
	repo "github.com/therobertcrocker/wayfinder/internal/machinations/repositories"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
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

// BulkAddAssets adds all of the assets to the storage layer
func (as *AssetStore) BulkAddAssets(filepath, fileType string) error {

	// Load the assets from the file
	switch fileType {
	case "csv":
		assets, err := loader.LoadAssetsFromCSV(filepath)
		if err != nil {
			return fmt.Errorf("failed to load assets from csv: %w", err)
		}

		// Add the assets to the data store
		for _, asset := range assets {

			// Create the asset data
			assetData := &domain.Asset{
				ID: util.NewAssetId(asset.Attribute, asset.Class, asset.Type),
				Info: domain.AssetInfo{
					Class:     asset.Class,
					Type:      asset.Type,
					Attribute: domain.ParseAttribute(asset.Attribute),
					TechLevel: domain.TechLevel(asset.TechLevel),
					Qualities: asset.Qualities(),
				},
				Action: domain.AssetAction{
					Description: asset.ActionDescription,
				},
				Cost: domain.AssetCost{
					Credits:   asset.Credits,
					Upkeep:    asset.Upkeep,
					Threshold: asset.AttributeThreshold,
				},
			}

			// Add the attack data if it exists
			if asset.TargetAttribute != "" {
				assetData.Attack = domain.Attack{
					TargetAttribute: domain.ParseAttribute(asset.TargetAttribute),
					UseAttribute:    domain.ParseAttribute(asset.UseAttribute),
					Damage: domain.Damage{
						Count:    asset.DamageCount,
						Die:      asset.DamageDie,
						Modifier: asset.DamageMod,
					},
				}
			}

			// Add the counter attack data if it exists
			if asset.CounterAttackCount != 0 {
				assetData.Counter = domain.Damage{
					Count:    asset.CounterAttackCount,
					Die:      asset.CounterAttackDie,
					Modifier: asset.CounterAttackMod,
				}
			}

			logging.Log.WithFields(log.Fields{"ID": assetData.ID}).Debug("Sucessfully built asset data")

			// Add the asset to the asset pool
			switch assetData.Info.Attribute {
			case domain.Cunning:
				as.AssetPool.Cunning[assetData.ID] = assetData
			case domain.Wealth:
				as.AssetPool.Wealth[assetData.ID] = assetData
			case domain.Force:
				as.AssetPool.Force[assetData.ID] = assetData
			}

		}
	default:
		return fmt.Errorf("file type %s not supported", fileType)

	}

	// Add the assets to the storage layer
	if err := as.AssetRepo.BulkAddAssets(as.AssetPool); err != nil {
		return fmt.Errorf("failed to bulk add assets: %w", err)
	}

	return nil
}
