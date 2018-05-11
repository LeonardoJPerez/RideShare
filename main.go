package main

import (
	"os"

	"github.com/RideShare-Server/configuration"
	"github.com/RideShare-Server/handlers/requestTypes"
	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/services/aws"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	configuration.SetupEnv()
	configuration.SetupCORS(e)

	database := configuration.InitializeDBConnection()
	configuration.SetupRouter(e, database)
	log.InitLog()

	startServer(e)
}

func TestLogin() {
	var authService = aws.NewCognitoService()
	u2 := ""
	password := ""
	authResult, err := authService.Authenticate(u2, password)
	if err != nil {
		panic(err)
	}

	err = authService.ValidateToken(authResult.AccessToken)
	if err != nil {
		panic(err)
	}
}

func startServer(e *echo.Echo) {
	e.Validator = &requestTypes.CustomValidator{
		Validator: validator.New(),
	}

	serverPort := os.Getenv("APP_PORT")
	if serverPort == "" {
		serverPort = ":8888"
	}

	e.Logger.Fatal(e.Start(serverPort))
}
