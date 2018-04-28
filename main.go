package main

import (
	"os"

	"github.com/RideShare-Server/aws"
	"github.com/RideShare-Server/configuration"
	"github.com/RideShare-Server/handlers/requests"
	"github.com/RideShare-Server/log"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
)

var awsService = aws.New()

func main() {
	e := echo.New()

	configuration.SetupEnv()
	configuration.SetupCORS(e)

	database := configuration.InitializeDBConnection(awsService)
	configuration.SetupRouter(e, database)
	log.InitLog()

	startServer(e)
}

func TestLogin() {
	u2 := ""
	password := ""
	authResult, err := awsService.Authenticate(u2, password)
	if err != nil {
		panic(err)
	}

	err = awsService.ValidateToken(authResult.AccessToken)
	if err != nil {
		panic(err)
	}
}

func startServer(e *echo.Echo) {
	e.Validator = &requests.CustomValidator{
		Validator: validator.New(),
	}

	serverPort := os.Getenv("APP_PORT")
	if serverPort == "" {
		serverPort = ":8888"
	}

	e.Logger.Fatal(e.Start(serverPort))
}
