package models

import (
	"github.com/jinzhu/gorm"
)

type Motorcycle struct {
	gorm.Model
	Displacement string `json:"displacement"`
	EngineHP     string `json:"engine_hp"`
	Image        string `json:"image"`
	Make         string `json:"make"`
	MakeID       int64  `json:"make_id"`
	ModelID      int64  `json:"model_id"`
	ModelName    string `json:"model"`
	Nickname     string `json:"nickname"`
	Thumbnail    string `json:"thumbnail"`
	VIN          string `json:"vin"`
	Year         int    `json:"year"`
}

// CollectionName collection name.
func (m Motorcycle) CollectionName() string {
	return "motorcycles"
}
