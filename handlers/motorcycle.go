package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/RideShare-Server/stores"
	"github.com/juju/errors"
	"github.com/labstack/echo"
)

// MotorcycleHandler :
type MotorcycleHandler struct {
	Store *stores.MotorcycleStore
}

// NewMotorcycleHandler :
func NewMotorcycleHandler() *MotorcycleHandler {
	c := new(MotorcycleHandler)
	return c
}

// GetMotorcycle :
func (c *MotorcycleHandler) GetMotorcycle(ctx echo.Context) error {
	id64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return errors.Trace(err)
	}

	return ctx.JSON(http.StatusOK, fmt.Sprintf("message: %d", id64))
}
