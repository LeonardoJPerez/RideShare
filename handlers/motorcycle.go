package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/discovry/streamfinderv3-api/store"
	"github.com/juju/errors"
	"github.com/labstack/echo"
)

type MotorcycleController struct {
	Store *store.ChannelStore
}

func NewChannelController() *MotorcycleController {
	c := new(MotorcycleController)
	return c
}

func (c *MotorcycleController) GetMotorcycle(ctx echo.Context) error {
	id64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return errors.Trace(err)
	}

	return ctx.JSON(http.StatusOK, fmt.Sprintf("message: %d", id64))
}
