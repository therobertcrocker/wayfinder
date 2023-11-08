package core

import (
	machinations "github.com/therobertcrocker/wayfinder/internal/machinations/data_stores"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
	"github.com/therobertcrocker/wayfinder/internal/util"
)

type Machinations struct {
	AssetStore   *machinations.AssetStore
	FactionStore *machinations.FactionStore
}

type Core struct {
	Config       *util.Config
	Storage      db.StorageManager
	Machinations *Machinations
}

func NewCore(config *util.Config) *Core {
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
