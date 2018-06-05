package configuration

import (
	"net/http"

	"github.com/RideShare-Server/handlers"
	"github.com/RideShare-Server/handlers/auth"
	"github.com/RideShare-Server/handlers/requestTypes"
	"github.com/RideShare-Server/models"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// Initialize AuthBoss decorator.
	ab := auth.SetupAuth()

	// Auth
	h := echo.WrapHandler(ab.NewRouter())
	e.GET("/auth/oauth2/google", h)
	e.GET("/auth/oauth2/callback/google", h)
	e.GET("/auth/oauth2/logout", h)

	// // Cognito User access.
	// authProvider := aws.NewCognitoService()
	// authHandler := handlers.NewAuthHandler(authProvider)
	// authRoutes := e.Group("/auth")
	// authRoutes.POST("", authHandler.Login)
	// authRoutes.POST("/validate", authHandler.ValidateToken)
	// authRoutes.POST("/change", authHandler.ChangePassword)

	// Motorcycle Handlers
	motorcycleHandler := handlers.NewMotorcycleHandler()
	motorcycleRoutes := e.Group("/bike")

	mw := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "header:Authorization",
	})

	motorcycleRoutes.GET("/u/:user_id", motorcycleHandler.GetByUser, mw)
	attachCRUDS(motorcycleRoutes, &models.Motorcycle{})

	// Ride Handlers
	ridesHandler := handlers.NewRideHandler()
	ridesRoutes := e.Group("/rides")
	ridesRoutes.POST("/dates", ridesHandler.GetRidesByDate)
	attachCRUDS(ridesRoutes, &models.Ride{})

	e.GET("/auth/google", redirectHandler)
	e.GET("/auth/google/callback", callbackHandler)
}

// hookCRUDS :
func attachCRUDS(routeGroup *echo.Group, model interface{}) {
	baseHandler := handlers.Base{}

	routeGroup.POST("", baseHandler.GetInsertHandler(model))
	routeGroup.GET("/:id", baseHandler.GetFetchHandler(model))
	routeGroup.PUT("", baseHandler.GetUpdateHandler(model))
	routeGroup.DELETE("/:id", baseHandler.GetDeleteHandler(model))
}

func redirectHandler(ctx echo.Context) error {
	return nil
	// authURL, err := gocial.New().
	// 	Driver("google"). // Set provider
	// 	Redirect()

	// if err != nil {
	// 	return ctx.JSON(http.StatusBadRequest, err)
	// }

	// return ctx.JSON(http.StatusOK, authURL)
	//c.Redirect(http.StatusFound, authURL) // Redirect with 302 HTTP code
}

func callbackHandler(ctx echo.Context) error {
	return nil
	// code := ctx.QueryParam("code")
	// state := ctx.QueryParam("state")

	// // Handle callback and check for errors
	// user, token, err := gocial.Handle(state, code)
	// if err != nil {
	// 	return ctx.JSON(http.StatusBadRequest, err)
	// }

	// // Print in terminal user information
	// fmt.Printf("%#v", token)
	// fmt.Printf("%#v", user)

	// // If no errors, show provider name
	// return ctx.JSON(http.StatusOK, user)
}
