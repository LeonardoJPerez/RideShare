package configuration

import (
	"net/http"

	"github.com/RideShare-Server/handlers"
	"github.com/RideShare-Server/handlers/requestTypes"
	"github.com/RideShare-Server/models"
	"github.com/RideShare-Server/services/aws"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
)

// SetupRouter inserts the application routes into the Echo context.
func SetupRouter(e *echo.Echo) {
	e.Validator = &requestTypes.CustomValidator{
		Validator: validator.New(),
	}

	// Base Route
	e.GET("/hb", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is up!")
	})

	// // Initialize AuthBoss decorator.
	// ab := auth.SetupAuth(db)

	// // Initialize db connection for middleware to use.
	// auth.SetupMiddleware(db)

	// // Auth
	// h := echo.WrapHandler(ab.NewRouter())
	// e.GET("/auth/oauth2/google", h)
	// e.GET("/auth/oauth2/callback/google", h)
	// e.GET("/auth/oauth2/logout", h)

	// Cognito User access.
	authProvider := aws.NewCognitoService()
	authHandler := handlers.NewAuthHandler(authProvider)
	authRoutes := e.Group("/auth")
	authRoutes.POST("", authHandler.Login)
	authRoutes.POST("/validate", authHandler.ValidateToken)
	authRoutes.POST("/change", authHandler.ChangePassword)

	motorcycleHandler := handlers.NewMotorcycleHandler()
	motorcycleRoutes := e.Group("/bike")

	model := &models.Motorcycle{}
	motorcycleRoutes.POST("", motorcycleHandler.GetInsertHandler(model))
	motorcycleRoutes.PUT("", motorcycleHandler.GetUpdateHandler(model))
	motorcycleRoutes.DELETE("/:id", motorcycleHandler.GetDeleteHandler(model))
	motorcycleRoutes.GET("/:id", motorcycleHandler.GetFetchHandler(model))
	motorcycleRoutes.GET("/u/:user_id", motorcycleHandler.GetByUser)

}
