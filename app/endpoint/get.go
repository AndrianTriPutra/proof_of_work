package endpoint

import (
	"atp/payment/pkg/utils/echos/util"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h handler) GetbyKey(c echo.Context) error {
	ctx := c.Request().Context()
	key := c.FormValue("key")
	if key == "" {
		return util.CustomError{
			ErrorType: util.ErrBadRequest,
			Message:   "The given data was invalid",
			Cause:     "key is required",
		}
	}

	data, err := h.repo.FindTransaction(ctx, key)
	if errors.Is(err, util.ErrorNotFound) {
		return util.CustomError{
			ErrorType: util.ErrBadRequest,
			Message:   "can't find key",
			Cause:     "key not found",
		}
	} else if err != nil {
		return util.CustomError{
			ErrorType: util.ErrInternalServer,
			Message:   "can't find key",
			Cause:     "something wrong",
		}
	}

	response := util.WrapSuccessResponse("success", data)
	return c.JSON(http.StatusOK, response)
}

func (h handler) GetALL(c echo.Context) error {
	ctx := c.Request().Context()
	data, err := h.repo.GetALLTrans(ctx)
	if errors.Is(err, util.ErrorNotFound) {
		return util.CustomError{
			ErrorType: util.ErrBadRequest,
			Message:   "can't find key",
			Cause:     "data not found",
		}
	} else if err != nil {
		return util.CustomError{
			ErrorType: util.ErrInternalServer,
			Message:   "can't find key",
			Cause:     "something wrong",
		}
	}

	response := util.WrapSuccessResponse("success", data)
	return c.JSON(http.StatusOK, response)
}
