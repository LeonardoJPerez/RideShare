package models

import (
	"github.com/RideShare-Server/models/enums"
	"github.com/jinzhu/gorm"
)

type (
	User struct {
		gorm.Model
		City        string
		DateofBirth string
		DislayName  string
		Email       string
		Gender      string
		Logo        string // S3 resource url
		Name        string
		SkillLevel  enums.SkillLevel
		AddressID   uint
		Locale      string
	}
)
