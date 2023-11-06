package core

import (
	"github.com/therobertcrocker/wayfinder/internal/machinations"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
	"github.com/therobertcrocker/wayfinder/internal/util"
)

type Core struct {
	Config       *util.Config
	Storage      db.StorageManager
	Machinations *machinations.Machinations
}

func NewCore(config *util.Config) *Core {
	storage := db.NewBboltStorageManager(config.Database.Path, config.Database.Buckets)
	machinations := machinations.NewMachinations(storage)

	return &Core{
		Config:       config,
		Storage:      storage,
		Machinations: machinations,
	}

}
