package loader

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
	"github.com/therobertcrocker/wayfinder/internal/common/logging"
	"github.com/therobertcrocker/wayfinder/internal/machinations/domain"
)

type AssetCSV struct {
	Class              string `csv:"class"`
	Attribute          string `csv:"attribute"`
	Type               string `csv:"type"`
	ActionDescription  string `csv:"action_description"`
	Credits            int    `csv:"credits"`
	Upkeep             int    `csv:"upkeep"`
	TechLevel          int    `csv:"tech_level"`
	AttributeThreshold int    `csv:"attribute_threshold"`
	TargetAttribute    string `csv:"target_attribute"`
	UseAttribute       string `csv:"use_attribute"`
	DamageCount        int    `csv:"damage_count"`
	DamageDie          int    `csv:"damage_die"`
	DamageMod          int    `csv:"damage_mod"`
	CounterAttackCount int    `csv:"counter_attack_count"`
	CounterAttackDie   int    `csv:"counter_attack_die"`
	CounterAttackMod   int    `csv:"counter_attack_mod"`
	Quality1           string `csv:"quality_1"`
	Quality2           string `csv:"quality_2"`
	Quality3           string `csv:"quality_3"`
}

func (a *AssetCSV) Qualities() []string {
	qualities := []string{}
	if a.Quality1 != "" {
		qualities = append(qualities, a.Quality1)
	}
	if a.Quality2 != "" {
		qualities = append(qualities, a.Quality2)
	}
	if a.Quality3 != "" {
		qualities = append(qualities, a.Quality3)
	}

	return qualities
}

func LoadAssetsFromCSV(filepath string) ([]AssetCSV, error) {
	// Open the file

	logging.Log.WithFields(log.Fields{"filepath": filepath}).Debug("Loading assets from file")

	f, err := os.Open(filepath)
	if err != nil {
		logging.Log.WithFields(log.Fields{"filepath": filepath}).Errorf("failed to open file: %v", err)
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	defer f.Close()

	// Parse the file
	assets := []AssetCSV{}
	if err := gocsv.UnmarshalFile(f, &assets); err != nil {
		logging.Log.WithFields(log.Fields{"filepath": filepath}).Errorf("failed to unmarshal assets from file: %v", err)
		return nil, fmt.Errorf("failed to unmarshal assets from file: %w", err)
	}

	if err := validateAssets(assets); err != nil {
		return nil, err
	}

	return assets, nil

}

func validateAssets(assets []AssetCSV) error {

	logging.Log.WithFields(log.Fields{"count": len(assets)}).Debug("Validating assets")

	for i, assetData := range assets {
		if err := validateAsset(assetData); err != nil {
			logging.Log.WithFields(log.Fields{"row": i, "class": assetData.Class}).Errorf("failed to validate asset: %v", err)
			return fmt.Errorf("failed to validate asset at row %d | %s: %w", i, assetData.Class, err)
		}
	}

	return nil
}

func validateAsset(assetData AssetCSV) error {

	// Validate that class is not empty
	if assetData.Class == "" {
		return fmt.Errorf("class cannot be empty")
	}

	// Validate that type is not empty
	if assetData.Type == "" {
		return fmt.Errorf("type cannot be empty")
	}

	// validate that the Attribute is in the list of valid attributes
	if _, ok := domain.AttributeMap[assetData.Attribute]; !ok {
		return fmt.Errorf("attribute %s not valid", assetData.Attribute)
	}

	// validate that credits and upkeep are positive
	if assetData.Credits < 0 {
		return fmt.Errorf("credits cannot be negative")
	}

	if assetData.Upkeep < 0 {
		return fmt.Errorf("upkeep cannot be negative")
	}

	// validate that tech level is between 0 and the max tech level
	if _, ok := domain.TechLevelMap[assetData.TechLevel]; !ok {
		return fmt.Errorf("tech level %d not valid", assetData.TechLevel)
	}

	// validate that attribute threshold is positive
	if assetData.AttributeThreshold < 0 {
		return fmt.Errorf("attribute threshold cannot be negative")
	}

	// check if target attribute exists (its optional)
	if assetData.TargetAttribute != "" {

		// validate that the Target Attribute is in the list of valid attributes
		if _, ok := domain.AttributeMap[assetData.TargetAttribute]; !ok {
			return fmt.Errorf("target attribute %s not valid", assetData.TargetAttribute)
		}

		// validate that the Use Attribute is in the list of valid attributes
		if _, ok := domain.AttributeMap[assetData.UseAttribute]; !ok {
			return fmt.Errorf("use attribute %s not valid", assetData.UseAttribute)
		}

		// validate that damage count is positive
		if assetData.DamageCount < 0 {
			return fmt.Errorf("damage count cannot be negative")
		}

		// validate that damage die is positive
		if assetData.DamageDie < 0 {
			return fmt.Errorf("damage die cannot be negative")
		}

	}

	// check if counter attack exists (its optional)
	if assetData.CounterAttackCount != 0 {

		// validate that counter attack count is positive
		if assetData.CounterAttackCount < 0 {
			return fmt.Errorf("counter attack count cannot be negative")
		}

		// validate that counter attack die is positive
		if assetData.CounterAttackDie < 0 {
			return fmt.Errorf("counter attack die cannot be negative")
		}
	}

	return nil
}
