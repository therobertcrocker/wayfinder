package repositories

import (
	"errors"

	"github.com/therobertcrocker/wayfinder/internal/machinations/entities"
	db "github.com/therobertcrocker/wayfinder/internal/storage"
	"github.com/therobertcrocker/wayfinder/internal/util"
)

type FactionRepository interface {
	// Add Faction to the storage
	AddFaction(faction *entities.Faction) error

	// Get Faction from the storage
	GetFaction(id util.EntityID) (*entities.Faction, error)

	// Update Faction in the storage
	UpdateFaction(faction *entities.Faction) error

	// Delete Faction from the storage
	DeleteFaction(id util.EntityID) error
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

// Add Faction to the storage
func (r *FactionRepo) AddFaction(faction *entities.Faction) error {
	return errors.New("not implemented")
}

// Get Faction from the storage
func (r *FactionRepo) GetFaction(id util.EntityID) (*entities.Faction, error) {
	return nil, errors.New("not implemented")
}

// Update Faction in the storage
func (r *FactionRepo) UpdateFaction(faction *entities.Faction) error {
	return errors.New("not implemented")
}

// Delete Faction from the storage
func (r *FactionRepo) DeleteFaction(id util.EntityID) error {
	return errors.New("not implemented")
}
