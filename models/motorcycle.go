package models

import (
	"github.com/jinzhu/gorm"
)

type Motorcycle struct {
	gorm.Model
	Nickname  string
	BrandName string
	BrandID   uint
}

// CollectionName collection name.
func (m Motorcycle) CollectionName() string {
	return "motorcycles"
}
