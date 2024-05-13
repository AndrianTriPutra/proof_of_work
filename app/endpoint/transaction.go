package endpoint

import (
	"atp/payment/pkg/utils/domain"
	"atp/payment/pkg/utils/echos/util"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h handler) Transaction(c echo.Context) error {
	ctx := c.Request().Context()
	var data domain.Data
	err := c.Bind(&data)
	if err != nil {
		return util.CustomError{
			ErrorType: util.ErrBadRequest,
			Message:   "The given data was invalid",
			Cause:     "failed decode input",
		}
	}

	h.bc.GiveData(data)

	prev, _ := h.ucase.LatestBlock(ctx)

	var nonce uint
	log.Printf("version fix:%s", h.setting.Version)
	if h.setting.Version == "v1" {
		nonce = h.poW.PoW1(h.bc, prev.Key) //v1
	} else if h.setting.Version == "v2" {
		//v2
		pow := h.poW.NewProof(h.bc)
		dataX := new(domain.Data)
		dataX.From = data.From
		dataX.To = data.To
		dataX.IDR = data.IDR
		var dataY []*domain.Data
		dataY = append(dataY, dataX)
		nonce, _ = h.poW.Run(pow, prev.Key, dataY)
		valid := h.poW.Validate(pow, nonce, prev.Key, dataY)
		log.Printf("valid ? %s", strconv.FormatBool(valid))
	} else {
		nonce = h.poW.PoW3(h.setting.Version, h.bc, prev.Key) //v3
	}

	block := h.ucase.CreateBlock(ctx, nonce, h.bc, prev.Key)

	response := util.WrapSuccessResponse("success", block)
	return c.JSON(http.StatusOK, response)
}
