package handlers

import (
	"net/http"

	"github.com/RideShare-Server/stores"
	"github.com/juju/errors"
	"github.com/labstack/echo"
)

// RideHandler :
type RideHandler struct {
	Base
}

// NewRideHandler :
func NewRideHandler() *RideHandler {
	c := new(RideHandler)

	return c
}

/*
Methods
- GetRidesByLocation(startLocation *Address)
- GetRidesStartingBy(from, to string)
- GetRidesByOwnerID(ownerID uint)
- GetRidesByTags(tags []string)
*/

func (c *RideHandler) GetRidesByDate(ctx echo.Context) error {
	from := ctx.Param("from")
	to := ctx.Param("to")
	if from == "" {
		return ctx.JSON(http.StatusBadRequest, errors.Errorf("from datestamp cannot be empty"))
	}

	result, err := stores.Rides.GetStartingBy(from, to)
	if err != nil {
		return errors.Trace(err)
	}

	return ctx.JSON(http.StatusOK, result)
}
