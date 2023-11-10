package domain

import (
	"strings"
)

// --------------- Affiliation ---------------

// Affiliation is a custom type for the affiliation of a Faction.
type Affiliation int

const (
	Political Affiliation = iota
	Religious
	Paramilitary
	Corporate
)

// Affilition Map allows for parsing of Affiliation from a string - and for validating if an affiliation is valid.
var AffiliationMap = map[string]Affiliation{
	"political":    Political,
	"religious":    Religious,
	"paramilitary": Paramilitary,
	"corporate":    Corporate,
}

// ParseAffiliation parses a string into an Affiliation.
func ParseAffiliation(affiliation string) Affiliation {
	// normalize the affiliation string
	affiliation = strings.ToLower(affiliation)

	return AffiliationMap[affiliation]
}

// String returns the string representation of an Affiliation.
func (a Affiliation) String() string {
	return [...]string{"Political", "Religious", "Paramilitary", "Corporate"}[a]
}

// --------------- Size ---------------

// Size is a custom type for the size of a Faction.
type Size int

const (
	Small Size = iota
	Medium
	Large
)

// Size Map allows for parsing of Size from a string - and for validating if a size is valid.
var SizeMap = map[string]Size{
	"small":  Small,
	"medium": Medium,
	"large":  Large,
}

// ParseSize parses a string into a Size.
func ParseSize(size string) Size {
	// normalize the size string
	size = strings.ToLower(size)

	return SizeMap[size]

}

// String returns the string representation of a Size.
func (s Size) String() string {
	return [...]string{"Small", "Medium", "Large"}[s]
}

// --------------- TechLevel ---------------

// TechLevel is a custom type for the tech level of a Faction.
type TechLevel int

const (
	PreArcane TechLevel = iota
	Arcane
	ArcanoTech
	AetherAware
	AetherTech
)

var TechLevelMap = map[int]TechLevel{
	0: PreArcane,
	1: Arcane,
	2: ArcanoTech,
	3: AetherAware,
	4: AetherTech,
}

// String returns the string representation of a TechLevel.
func (tl TechLevel) String() string {
	return [...]string{"PreArcane", "Arcane", "ArcanoTech", "AetherAware", "AetherTech"}[tl]
}

// --------------- Attribute ---------------

// Attribute is a custom type for the attribute of an Asset.
type Attribute int

const (
	Cunning Attribute = iota
	Wealth
	Force
)

// Attribute Map allows for parsing of Attribute from a string - and for validating if an attribute is valid.
var AttributeMap = map[string]Attribute{
	"cunning": Cunning,
	"wealth":  Wealth,
	"force":   Force,
}

// ParseAttribute parses a string into an Attribute.
func ParseAttribute(attribute string) Attribute {
	// normalize the attribute string
	attribute = strings.ToLower(attribute)

	return AttributeMap[attribute]

}

// String returns the string representation of an Attribute.
func (a Attribute) String() string {
	return [...]string{"Cunning", "Wealth", "Force"}[a]
}
