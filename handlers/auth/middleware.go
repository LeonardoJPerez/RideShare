package auth

import (
	"fmt"
	"strconv"

	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/stores"
	"github.com/jinzhu/gorm"
	"github.com/juju/errors"
	"github.com/labstack/echo"
)

// ProtectRoute :
func ProtectRoute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		w := c.Response().Writer
		r := c.Request()
		u, err := ab.CurrentUser(w, r)
		if err != nil {
			log.Error(log.AuthMiddlewareTopic, err)
			return echo.ErrUnauthorized
		} else if u == nil {
			log.Error(log.AuthMiddlewareTopic, fmt.Errorf("Unauthorized user trying to access: %v", r.URL.Path))
			return echo.ErrUnauthorized
		}

		c.SetParamNames("userEmail")
		c.SetParamValues(u.(*stores.AuthResult).Email)

		return next(c)
	}
}

var (
	userSessionStore *stores.SessionStore
)

// ValidateSession :
func ValidateSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.ParseUint(c.QueryParam("uid"), 0, 64)
		if err != nil {
			log.Error(log.SessionMiddlewareTopic, err)
			return echo.ErrUnauthorized
		}

		sessionID := c.QueryParam("sid")
		if sessionID == "" {
			log.Error(log.SessionMiddlewareTopic, errors.New("session id missing from request"))
			return echo.ErrUnauthorized
		}

		// Validate session is active.
		isSessionValid, err := userSessionStore.Validate(uint(userID), sessionID)
		if err != nil || !isSessionValid {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

// SetupMiddleware :
func SetupMiddleware(database *gorm.DB) {
	userSessionStore = stores.NewSessionStore(database)
}
