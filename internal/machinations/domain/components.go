package domain

import "github.com/therobertcrocker/wayfinder/internal/common/util"

// ---------------------- Components ----------------------

type FactionInfo struct {
	Affiliation Affiliation   `json:"affiliation"`
	Size        Size          `json:"size"`
	TechLevel   TechLevel     `json:"tech_level"`
	HomeBase    util.EntityID `json:"home_base"`
}

type FactionAttributes struct {
	Cunning int `json:"cunning"`
	Wealth  int `json:"wealth"`
	Force   int `json:"force"`
	XP      int `json:"xp"`
}

type FactionWealth struct {
	Credits int `json:"credits"`
	Income  int `json:"income"`
}

type AssetInfo struct {
	Class     string    `json:"class"`
	Type      string    `json:"type"`
	Attribute Attribute `json:"attribute"`
	TechLevel TechLevel `json:"tech_level"`
	Qualities []string  `json:"qualities"`
}

type AssetAction struct {
	Description string `json:"description"`
}

type AssetCost struct {
	Credits   int `json:"credits"`
	Upkeep    int `json:"upkeep"`
	Threshold int `json:"threshold"`
}

type Attack struct {
	TargetAttribute Attribute `json:"target_attribute"`
	UseAttribute    Attribute `json:"use_attribute"`
	Damage          Damage    `json:"damage"`
}

type Damage struct {
	Count    int `json:"count"`
	Die      int `json:"die"`
	Modifier int `json:"modifier"`
}

type Health struct {
	Current int `json:"current"`
	Max     int `json:"max"`
}
