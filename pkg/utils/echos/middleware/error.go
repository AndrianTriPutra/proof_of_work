package middleware

import (
	"atp/payment/pkg/utils/echos/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

var ErrorMiddleware = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return err
		}

		customErr, ok := err.(util.CustomError)
		if !ok {
			resp := util.WrapErrorResponse("internal server error", nil)
			return c.JSON(http.StatusInternalServerError, resp)
		}
		switch customErr.ErrorType {
		case util.ErrInternalServer:
			resp := util.WrapErrorResponse(customErr.Message, customErr.Cause)
			return c.JSON(http.StatusInternalServerError, resp)
		default:
			resp := util.WrapErrorResponse(customErr.Message, customErr.Cause)
			return c.JSON(util.GetHttpStatusCode(customErr.ErrorType), resp)
		}
	}
}
