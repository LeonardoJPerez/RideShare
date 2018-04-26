package models

import (
	"github.com/RideShare-Server/models/enums"
	"github.com/jinzhu/gorm"
)

type (
	Address struct {
		ID      uint
		Line1   string
		Line2   string
		City    string
		State   string
		Country string
	}

	Comment struct {
		ID          uint
		UserID      string
		DateCreated string
		Text        string
	}

	RouteMarker struct {
		ID        uint
		Latitude  string
		Longitude string
		AddressID uint
	}

	PitStop struct {
		RouteMarker
		Name string
		Icon string
	}

	Ride struct {
		gorm.Model
		StartLocation  string
		FinishLocation string
		When           string
		Description    string
		HostsIDs       []string
		Route          []*RouteMarker
		Comments       []*Comment
		Logo           string
		Image          string
		Attending      uint
		IsPrivate      bool
		Tags           []string
		SkillLevel     enums.SkillLevel
		RidingStyle    enums.RidingStyle
		RideType       enums.RideType
		MakeBrand      string
	}
)
