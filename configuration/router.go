package configuration

import (
	"net/http"

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
}
