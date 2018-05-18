package handlers

import (
	"net/http"
	"strconv"

	"github.com/RideShare-Server/stores"

	"github.com/juju/errors"
	"github.com/labstack/echo"
)

// MotorcycleHandler :
type MotorcycleHandler struct {
	Base
}

// NewMotorcycleHandler :
func NewMotorcycleHandler() *MotorcycleHandler {
	c := new(MotorcycleHandler)

	return c
}

// GetByUser retrieves all motorcycles belonging to a UserID.
func (c *MotorcycleHandler) GetByUser(ctx echo.Context) error {
	id64, err := strconv.ParseUint(ctx.Param("user_id"), 10, 64)
	if err != nil {
		return errors.Trace(err)
	}

	result, err := stores.Motorcycles.GetByUser(uint(id64))
	if err != nil {
		return errors.Trace(err)
	}

	return ctx.JSON(http.StatusOK, result)
}
