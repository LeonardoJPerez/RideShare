package configuration

import (
	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/stores"
	"github.com/labstack/echo"
)

// Inject bootstraps the application wide dependencies i.e. Db connection, Routes, Environment variables, etc.
func Inject(e *echo.Echo) {
	SetupEnv()

	database := InitializeDBConnection()
	stores.Init(database)

	SetupCORS(e)
	SetupRouter(e, database)

	log.InitLog()
}
