package main

import (
	"os"

	"github.com/RideShare-Server/configuration"
	"github.com/RideShare-Server/services/aws"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	configuration.Inject(e)

	serverPort := os.Getenv("APP_PORT")
	if serverPort == "" {
		serverPort = ":8888"
	}

	e.Logger.Fatal(e.Start(serverPort))
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
