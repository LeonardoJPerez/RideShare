package configuration

import (
	"fmt"
	"os"

	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/utils"

	"github.com/joho/godotenv"
	"github.com/juju/errors"
)

// SetupEnv method checks if application is being executed in a Production environment.
// If application is NOT running in production the method will look for a .env file and load the values from it.
func SetupEnv() {
	_, inProduction := os.LookupEnv(utils.Environment)
	if inProduction {
		return
	}

	fmt.Println("Injecting local ENV variables...")
	err := godotenv.Load()
	if err != nil {
		err = fmt.Errorf("Error loading .env file: %v", errors.Cause(err))
		log.Error(log.DbConnectionTopic, err)
	}
}
