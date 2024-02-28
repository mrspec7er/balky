package logger

import (
	"net/http"

	"github.com/mrspec7er/balky/app/utils"
)

type LoggerController struct {
	service  LoggerService
	response utils.Response
}

func (c *LoggerController) FindAll(w http.ResponseWriter, r *http.Request) {
	result, status, err := c.service.FindMany()

	if err != nil {
		c.response.InternalServerErrorHandler(w, status, err)
	}

	c.response.GetSuccessResponse(w, nil, result, nil)
}
