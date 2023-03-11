// Package generator implements generators like UUIDGenerator
package generator

import (
	"github.com/google/uuid"
)

// UUIDGenerator wrapper around google/uuid that allows this type to satisfy interfaces
type UUIDGenerator struct{}

// NewID returns a new UUID
func (g UUIDGenerator) NewID() string {
	return uuid.NewString()
}
