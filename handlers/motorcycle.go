package handlers

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/RideShare-Server/models"
	"github.com/RideShare-Server/stores"
	"github.com/juju/errors"
	"github.com/labstack/echo"
)

// MotorcycleHandler :
type MotorcycleHandler struct {
	Store *stores.MotorcycleStore
}

// NewMotorcycleHandler :
func NewMotorcycleHandler(db *gorm.DB) *MotorcycleHandler {
	c := new(MotorcycleHandler)

	c.Store = stores.NewMotorcycleStore(db)

	return c
}

// GetMotorcycleByID retrieves a motorcycle by its ID.
func (c *MotorcycleHandler) GetMotorcycleByID(ctx echo.Context) error {
	id64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return errors.Trace(err)
	}

	result, err := c.Store.GetByID(uint(id64))
	if err != nil {
		return errors.Trace(err)
	}

	return ctx.JSON(http.StatusOK, result)
}

// GetMotorcycleByUser retrieves all motorcycles belonging to a UserID.
func (c *MotorcycleHandler) GetMotorcycleByUser(ctx echo.Context) error {
	id64, err := strconv.ParseUint(ctx.Param("user_id"), 10, 64)
	if err != nil {
		return errors.Trace(err)
	}

	result, err := c.Store.GetByUser(uint(id64))
	if err != nil {
		return errors.Trace(err)
	}

	return ctx.JSON(http.StatusOK, result)
}

// InsertMotorcycle creates a new Motorcycle record.
func (c *MotorcycleHandler) InsertMotorcycle(ctx echo.Context) error {
	m := &models.Motorcycle{}
	if err := ctx.Bind(m); err != nil {
		return errors.Trace(err)
	}

	if err := ctx.Validate(m); err != nil {
		return errors.Trace(err)
	}

	result, err := c.Store.Insert(m)
	if err != nil {
		return errors.Trace(err)
	}

	return ctx.JSON(http.StatusOK, result)
}
