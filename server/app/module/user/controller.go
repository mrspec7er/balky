package user

import (
	"net/http"

	"github.com/mrspec7er/balky/app/utils"
)

type UserController struct {
	service  UserService
	response utils.Response
}

func (c *UserController) FindAllController(w http.ResponseWriter, r *http.Request) {
	result, status, err := c.service.FindManyService()

	if err != nil {
		c.response.InternalServerErrorHandler(w, status, err)
	}

	c.response.GetSuccessResponse(w, nil, result, nil)
}
