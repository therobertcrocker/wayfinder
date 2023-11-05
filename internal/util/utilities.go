package util

import "fmt"

type EntityID string

func NewEntityID(entityType, name string) EntityID {
	return EntityID(fmt.Sprintf("%s_%s", entityType, name))
}
