package main

import (
	"errors"
	"os"

	"github.com/RideShare-Server/aws"
	"github.com/RideShare-Server/configuration"
	"github.com/RideShare-Server/db"
	"github.com/RideShare-Server/handlers/requests"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
)

var awsService = aws.New()

func main() {
	//e := echo.New()

	// 1. Load our env if needed.
	configuration.SetupEnv()

	// 2. Setup our database connection.

	// 3. Add cors middlewear.
	//configuration.SetupCors(e)

	// 4. Route our requests.
	//configuration.SetupRouter(e, database)

	// 7. Setup our logger.
	//log.InitLog()

	// 8. Fire up the server.
	//startServer(e)
	// a := nhtsa.NewService()
	// models, err := a.GetModels("474")
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	// u1 := "find.leonardo@gmail.com"
	u2 := "<EMAIL>"
	password := "<PASSWORD>"
	authResult, err := awsService.Authenticate(u2, password)
	if err != nil {
		panic(err)
	}

	err = awsService.ValidateToken(authResult.AccessToken)
	if err != nil {
		panic(err)
	}
}

func TestDBConnection() {
	connectionString, err := awsService.GetRDSConnectionString()
	if err != nil {
		panic(err)
	}

	dbContext, err := db.Connect(connectionString)
	if err != nil {
		panic(err)
		return
	}

	if dbContext == nil {
		panic(errors.New("DB context is null"))
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
