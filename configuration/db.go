package configuration

import (
	"github.com/RideShare-Server/db"
	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/utils"
	"github.com/jinzhu/gorm"
)

// InitializeDBConnection returns a instance of a DB connection.
func InitializeDBConnection() *gorm.DB {
	dsn := utils.GetEnvVariable("CONNECTION_STRING")

	database, err := db.Connect(dsn)
	if err != nil {
		log.Error(log.DbConnectionTopic, err)
		return nil
	}

	// stats := database.DB().Stats()
	return database
}
