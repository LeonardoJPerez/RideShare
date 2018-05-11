package handlers

import (
	"net/http"

	"github.com/RideShare-Server/handlers/requestTypes"
	"github.com/RideShare-Server/services"

	"github.com/labstack/echo"
)

// AuthHandler :
type AuthHandler struct {
	authProvider services.AuthInterface
}

// NewAuthHandler :
func NewAuthHandler(authProvider services.AuthInterface) *AuthHandler {
	c := new(AuthHandler)
	c.authProvider = authProvider

	return c
}

// Login :
func (c *AuthHandler) Login(ctx echo.Context) error {
	request := new(requestTypes.LoginRequest)
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Authenticate and return token.
	res, err := c.authProvider.Authenticate(request.UserName, request.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
