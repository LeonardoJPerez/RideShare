package main

import (
	"os"

	"github.com/RideShare-Server/aws"
	"github.com/RideShare-Server/configuration"
	"github.com/RideShare-Server/db"
	"github.com/RideShare-Server/handlers/requests"
	"github.com/RideShare-Server/utils"

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

	TestDBConnection()
}

func TestDBConnection() {
	dsn, err := awsService.GetDns()
	if err != nil {
		panic(err)
	}

	utils.PrettyPrint(dsn)

	dbContext, err := db.Connect(dsn)
	if err != nil {
		panic(err)
		return
	}

	stats := dbContext.DB().Stats()
	utils.PrettyPrint(stats)
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
