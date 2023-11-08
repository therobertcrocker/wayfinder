package domain

import (
	"github.com/therobertcrocker/wayfinder/internal/util"
)

// ---------------------- Entities ----------------------

type Faction struct {
	ID         util.EntityID                 `json:"id"`
	Name       string                        `json:"name"`
	Info       FactionInfo                   `json:"info"`
	Attributes FactionAttributes             `json:"attributes"`
	Wealth     FactionWealth                 `json:"wealth"`
	Health     Health                        `json:"health"`
	Assets     map[util.EntityID]AssetRecord `json:"assets"`
}

type Asset struct {
	ID      util.EntityID `json:"id"`
	Info    AssetInfo     `json:"info"`
	Action  AssetAction   `json:"action"`
	Cost    AssetCost     `json:"cost"`
	Health  Health        `json:"health"`
	Attack  Attack        `json:"attack"`
	Counter Damage        `json:"counter"`
}

type AssetPool struct {
	Cunning map[util.EntityID]*Asset
	Wealth  map[util.EntityID]*Asset
	Force   map[util.EntityID]*Asset
}

type AssetRecord struct {
	ID       util.EntityID `json:"id"`
	Name     string        `json:"name"`
	Location util.EntityID `json:"location"`
	Data     Asset         `json:"data"`
}
