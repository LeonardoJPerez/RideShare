package utils

import (
	"github.com/RideShare-Server/models"
	"github.com/icrowley/fake"
)

// =======================
// Accounts Helpers
// =======================

// GetDummyMotorcycle :
func GetDummyMotorcycle() *models.Motorcycle {
	m := &models.Motorcycle{
		Nickname:     fake.FirstName(),
		Displacement: "471",
		Make:         "MV Agusta",
		ModelName:    "Brutale 675",
		ModelID:      "15",
		MakeID:       "16",
		EngineHP:     "125",
	}
	m.ID = 1

	return m
}
