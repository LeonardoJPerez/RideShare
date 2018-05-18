package configuration

import (
	"github.com/RideShare-Server/log"
	"github.com/labstack/echo"
)

// Inject bootstraps the application wide dependencies i.e. Db connection, Routes, Environment variables, etc.
func Inject(e *echo.Echo) {
	SetupEnv()
	SetupCORS(e)

	database := InitializeDBConnection()
	SetupRouter(e, database)

	log.InitLog()
}
