package configuration

import (
	"net/http"

	"github.com/RideShare-Server/handlers"
	"github.com/RideShare-Server/services/aws"
	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

// SetupRouter inserts the application routes into the Echo context.
func SetupRouter(e *echo.Echo, db *gorm.DB) {
	//-----------------------
	// Base Routes
	//-----------------------

	// Base Route
	e.GET("/hb", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is up!")
	})

	authProvider := aws.NewCognitoService()
	authHandler := handlers.NewAuthHandler(authProvider)
	authRoutes := e.Group("/auth")
	authRoutes.POST("", authHandler.Login)
	authRoutes.POST("/validate", authHandler.ValidateToken)
	authRoutes.POST("/change", authHandler.ChangePassword)

	motorcycleHandler := handlers.NewMotorcycleHandler(db)
	motorcycleRoutes := e.Group("/bike")
	motorcycleRoutes.POST("", motorcycleHandler.InsertMotorcycle)
	motorcycleRoutes.GET("/:id", motorcycleHandler.GetMotorcycleByID)
	motorcycleRoutes.GET("/u/:user_id", motorcycleHandler.GetMotorcycleByUser)

}
