package models

import (
	"github.com/RideShare-Server/models/enums"
	"github.com/jinzhu/gorm"
)

type (

	// Address :
	Address struct {
		ID      uint
		Line1   string
		Line2   string
		City    string
		State   string
		Country string
	}

	// Comment :
	Comment struct {
		ID          uint
		UserID      string
		DateCreated string
		Text        string
	}

	// RouteMarker :
	RouteMarker struct {
		ID        uint
		Latitude  string
		Longitude string
		AddressID uint
		Name      string
		Icon      string
		IsPitStop bool
	}

	// Ride :
	Ride struct {
		gorm.Model
		StartLocation  uint              `gorm:"index;not null" json:"start_location"`
		FinishLocation uint              `gorm:"index;not null" json:"finish_location"`
		When           string            `gorm:"index;not null" json:"when"`
		Description    string            `gorm:"not null" json:"description"`
		Hosts          []string          `gorm:"-" json:"hosts"`
		Route          []*RouteMarker    `gorm:"-" json:"route"`
		Comments       []*Comment        `gorm:"-" json:"comments"`
		Logo           string            `gorm:"" json:"logo"`
		Name           string            `gorm:"unique_index" json:"name"`
		Image          string            `gorm:"" json:"image"`
		IsPrivate      bool              `gorm:"index;not null" json:"is_private"`
		Tags           []string          `gorm:"-" json:"tags"`
		SkillLevel     enums.SkillLevel  `gorm:"index;not null" json:"skill_level"`
		RidingStyle    enums.RidingStyle `gorm:"index;not null" json:"riding_style"`
	}
)
