package configuration

import (
	"net/http"

	"github.com/RideShare-Server/handlers"
	"github.com/RideShare-Server/handlers/auth"
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

	// Initialize AuthBoss decorator.
	ab := auth.SetupAuth(db)

	// Initialize db connection for middleware to use.
	auth.SetupMiddleware(db)

	// Auth
	h := echo.WrapHandler(ab.NewRouter())
	e.GET("/auth/oauth2/google", h)
	e.GET("/auth/oauth2/callback/google", h)
	e.GET("/auth/oauth2/logout", h)

	// Cognito User access.
	authProvider := aws.NewCognitoService()
	authHandler := handlers.NewAuthHandler(authProvider)
	authRoutes := e.Group("/auth")
	authRoutes.POST("", authHandler.Login)
	authRoutes.POST("/validate", authHandler.ValidateToken)
	authRoutes.POST("/change", authHandler.ChangePassword)

	motorcycleHandler := handlers.NewMotorcycleHandler(db)
	motorcycleRoutes := e.Group("/bike")
	motorcycleRoutes.POST("", motorcycleHandler.Insert)
	motorcycleRoutes.DELETE("/:id", motorcycleHandler.Remove)
	motorcycleRoutes.GET("/:id", motorcycleHandler.GetByID)
	motorcycleRoutes.GET("/u/:user_id", motorcycleHandler.GetByUser)

}
