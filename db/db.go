package db

import (
	"fmt"
	"os"
	"time"

	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/models"
	"github.com/RideShare-Server/utils"

	"github.com/jinzhu/gorm"

	// Bringing in MySQL dialects for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/juju/errors"
)

const (
	dbDialect    = "mysql"
	databaseName = "BikeMeet"
)

var (
	dbContext *gorm.DB
)

// Close :
func Close() {
	if err := dbContext.Close(); err != nil {
		err = errors.Errorf("Error closing the %v DB: %v", databaseName, errors.Cause(err))
		log.Error(log.DbConnectionTopic, err)
	}

	log.Info(log.DbConnectionTopic, fmt.Sprintf("Connection to %v DB closed", databaseName))
}

// Connect opens a connection to the database and runs migrations.
// Connectionstring format should be i the form of <USER>:<PASS>@tcp(<URI>:<PORT>)/<DBNAME>
func Connect(connectionString string) (*gorm.DB, error) {
	if dbContext != nil {
		return dbContext, nil
	}

	if connectionString == "" {
		return nil, fmt.Errorf("Connection string cannot be empty\n Format: USER:PASS@tcp(URI:PORT)/DBNAME")
	}

	log.Info(log.DbConnectionTopic, fmt.Sprintf("Attempting database connection - %v", databaseName))

	var err error
	retries := 10
	for i := retries; i > 0; i-- {
		dbContext, err = gorm.Open(dbDialect, connectionString)
		if err != nil {
			log.Error(log.DbConnectionTopic, err)
			time.Sleep(1 * time.Second)
		} else {
			log.Info(log.DbConnectionTopic, fmt.Sprintf("Connected to the database - %v", databaseName))
			break
		}
	}

	if dbContext == nil {
		return nil, errors.New("Db context could not be initialized")
	}

	dbContext.SingularTable(true)

	// Run auto migrations.
	migrate, migrateSet := os.LookupEnv(utils.Migrate)
	if migrateSet && migrate == "TRUE" {
		runAutoMigrations()
	}

	return dbContext, err
}

// CheckNotFoundErr inspects error object from gorm for a not found error message.
func CheckNotFoundErr(err error) error {
	if err == nil {
		return nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil
	}
	return err
}

// Context :
func Context() gorm.DB {
	return *dbContext
}

// Migrate schema soft changes
func runAutoMigrations() {
	dbContext.AutoMigrate(
		&models.Address{},
		&models.Comment{},
		&models.Motorcycle{},
		&models.Ride{},
		&models.RouteMarker{},
		&models.User{},
	)
}
