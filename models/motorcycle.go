package models

import (
	"github.com/jinzhu/gorm"
)

type Motorcycle struct {
	gorm.Model
	UserID       uint   `json:"user_id" validate:"required"`
	Displacement string `json:"displacement"`
	EngineHP     string `json:"engine_hp"`
	Image        string `json:"image"`
	Make         string `json:"make" validate:"required"`
	MakeID       int64  `json:"make_id" validate:"required"`
	ModelID      int64  `json:"model_id"`
	ModelName    string `json:"model" validate:"required"`
	Nickname     string `json:"nickname"`
	Thumbnail    string `json:"thumbnail"`
	VIN          string `json:"vin"`
	Year         uint   `json:"year" validate:"required"`
}

// CollectionName collection name.
func (m Motorcycle) CollectionName() string {
	return "motorcycles"
}
