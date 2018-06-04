package auth

import (
	"fmt"

	"github.com/juju/errors"
	"github.com/labstack/echo"
)

// ProtectRoute :
func ProtectRoute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		r := c.Request()
		if u, err := ab.CurrentUser(w, r); err != nil {
			fmt.Printf("\nError fetching the current user: %v\n", errors.Cause(err))
			return echo.ErrUnauthorized
		} else if u == nil {
			fmt.Printf("\nUnauthorized user trying to access: %v\n", r.URL.Path)
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
