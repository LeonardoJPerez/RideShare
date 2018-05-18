package handlers

import (
	"net/http"
	"strconv"

	"github.com/RideShare-Server/stores"
	"github.com/juju/errors"
	"github.com/labstack/echo"
)

// Base :
type Base struct {
}

// GetInsertHandler creates a handler that Inserts a new Motorcycle record.
func (bh *Base) GetInsertHandler(m interface{}) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		if err := ctx.Bind(m); err != nil {
			return errors.Trace(err)
		}

		if err := ctx.Validate(m); err != nil {
			return errors.Trace(err)
		}

		result, err := stores.Base.Insert(m)
		if err != nil {
			return errors.Trace(err)
		}

		return ctx.JSON(http.StatusOK, result)
	}

}

// GetUpdateHandler creates a handler that Updates an existing Resource record.
func (bh *Base) GetUpdateHandler(m interface{}) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		kvp := map[string]interface{}{}
		if err := ctx.Bind(&kvp); err != nil {
			return errors.Trace(err)
		}

		result, err := stores.Base.Update(kvp, m)
		if err != nil {
			return errors.Trace(err)
		}

		return ctx.JSON(http.StatusOK, result)
	}
}

// GetDeleteHandler creates a handler that Removes a Resource record.
func (bh *Base) GetDeleteHandler(m interface{}) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		id64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			return errors.Trace(err)
		}

		result, err := stores.Base.Remove(uint(id64), m)
		if err != nil {
			return errors.Trace(err)
		}

		return ctx.JSON(http.StatusOK, result)
	}
}

// GetFetchHandler creates a handler that Gets a Resource record.
func (bh *Base) GetFetchHandler(m interface{}) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		id64, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			return errors.Trace(err)
		}

		result, err := stores.Base.Get(uint(id64), m)
		if err != nil {
			return errors.Trace(err)
		}

		return ctx.JSON(http.StatusOK, result)
	}
}
