package core

import (
	"github.com/therobertcrocker/wayfinder/internal/common/config"
	machinations "github.com/therobertcrocker/wayfinder/internal/machinations/data_stores"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
)

type Machinations struct {
	AssetStore   *machinations.AssetStore
	FactionStore *machinations.FactionStore
}

type Core struct {
	Config       *config.Config
	Storage      db.StorageManager
	Machinations *Machinations
}

func NewCore(config *config.Config) *Core {
	storage := db.NewBboltStorageManager(config.Database.Path, config.Database.Buckets)
	machinations := &Machinations{
		AssetStore:   machinations.NewAssetStore(storage),
		FactionStore: machinations.NewFactionStore(storage),
	}

	return &Core{
		Config:       config,
		Storage:      storage,
		Machinations: machinations,
	}

}
