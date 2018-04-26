package models

import "github.com/jinzhu/gorm"

type Garage struct {
	gorm.Model
	Name            string
	UserID          string
	PreferredBikeID string
	MotorcycleIDs   []uint
}

// CollectionName collection name.
func (m Garage) CollectionName() string {
	return "garages"
}
