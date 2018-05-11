package configuration

import (
	"github.com/RideShare-Server/db"
	"github.com/RideShare-Server/utils"
	"github.com/jinzhu/gorm"
)

func InitializeDBConnection() *gorm.DB {
	dsn := utils.GetEnvVariable("CONNECTION_STRING")

	database, err := db.Connect(dsn)
	if err != nil {
		panic(err)
		return nil
	}

	stats := database.DB().Stats()
	utils.PrettyPrint(stats)

	return database
}
