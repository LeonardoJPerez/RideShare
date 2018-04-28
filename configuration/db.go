package configuration

import (
	"github.com/RideShare-Server/aws"
	"github.com/RideShare-Server/db"
	"github.com/RideShare-Server/utils"
	"github.com/jinzhu/gorm"
)

func InitializeDBConnection(s *aws.Service) *gorm.DB {
	dsn, err := s.GetDns()
	if err != nil {
		panic(err)
	}

	database, err := db.Connect(dsn)
	if err != nil {
		panic(err)
		return nil
	}

	stats := database.DB().Stats()
	utils.PrettyPrint(stats)

	return database
}
