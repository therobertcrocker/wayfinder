package domain

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

func (a Attribute) String() string {
	return [...]string{"Cunning", "Wealth", "Force"}[a]
}
