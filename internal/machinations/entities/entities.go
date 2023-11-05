package entities

import (
	"github.com/therobertcrocker/wayfinder/internal/util"
)

// ---------------------- Data Types ----------------------
type Affiliation int

const (
	Political Affiliation = iota
	Religious
	Paramilitary
	Corporate
)

type Size int

const (
	Small Size = iota
	Medium
	Large
)

type TechLevel int

const (
	PreArcane TechLevel = iota
	Arcane
	ArcanoTech
	AetherAware
	AetherTech
)

type Attribute int

const (
	Cunning Attribute = iota
	Wealth
	Force
)

// ---------------------- Components ----------------------

type FactionData struct {
	Name        string
	Affiliation Affiliation
	Size        Size
	TechLevel   TechLevel
	HomeBase    util.EntityID
}

type FactionAttributes struct {
	Cunning int
	Wealth  int
	Force   int
	XP      int
}

type FactionWealth struct {
	Credits int
	Income  int
}

type AssetData struct {
	Name        string
	Attribute   Attribute
	Description string
	TechLevel   TechLevel
	Qualities   []string
	Location    util.EntityID
}

type AssetCost struct {
	Credits   int
	Upkeep    int
	Threshold int
}

type Attack struct {
	TargetAttribute Attribute
	UseAttribute    Attribute
	Damage          Damage
}

type Damage struct {
	Count    int
	Die      int
	Modifier int
}

type Health struct {
	Current int
	Max     int
}

// ---------------------- Entities ----------------------

type Faction struct {
	ID         util.EntityID
	Data       FactionData
	Attributes FactionAttributes
	Wealth     FactionWealth
	Health     Health
	Assets     map[util.EntityID]Asset
}

type Asset struct {
	ID      util.EntityID
	Data    AssetData
	Cost    AssetCost
	Health  Health
	Attack  Attack
	Counter Damage
}
