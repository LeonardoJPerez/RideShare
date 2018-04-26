package handlers

import (
	"net/http"

	"github.com/RideShare-Server/handlers/requests"
	"github.com/discovry/streamfinderv3-api/store"
	"github.com/labstack/echo"
)

type AuthController struct {
	Store *store.ChannelStore
}

func NewAuthController() *AuthController {
	c := new(AuthController)
	return c
}

func (c *MotorcycleController) Login(ctx echo.Context) error {
	request := new(requests.LoginRequest)
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Authenticate and return token.

	return ctx.JSON(http.StatusOK, "{success='true'}")
}
