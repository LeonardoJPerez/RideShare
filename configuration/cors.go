package configuration

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// SetupCORS configuration method for Echo instance.
func SetupCORS(e *echo.Echo) {
	config := middleware.DefaultCORSConfig
	config.AllowCredentials = true
	config.AllowOrigins = []string{
		"http://localhost:8881",
	}

	e.Use(middleware.CORSWithConfig(config))
}
