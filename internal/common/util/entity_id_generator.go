package util

import (
	"strings"
)

type EntityID string

// NewEntityID creates a new EntityID from the given fields
func NewEntityID(entityType string, fields ...string) EntityID {
	allFields := append([]string{entityType}, fields...)
	return EntityID(strings.Join(allFields, "_"))
}

// New AssetID enforces the correct format for an AssetID and calls NewEntityID
func NewAssetId(attribute, class, assetType string) EntityID {

	return NewEntityID("asset", attribute, class, assetType)
}
