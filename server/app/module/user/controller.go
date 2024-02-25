package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/utils"
)

type UserController struct {
	service  UserService
	response utils.Response
	publish  utils.Publisher
}

func (c *UserController) FindAllController(w http.ResponseWriter, r *http.Request) {
	result, status, err := c.service.FindManyService()

	if err != nil {
		c.response.InternalServerErrorHandler(w, status, err)
	}

	c.response.GetSuccessResponse(w, nil, result, nil)
}

func (c *UserController) CreateController(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
	}

	ctx := context.Background()
	err = c.publish.SendMessage(ctx, "user.create", data)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
	}

	c.response.SuccessMessageResponse(w, "create user with email: "+user.Email)
}
