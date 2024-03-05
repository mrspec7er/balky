package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/utility"
)

type UserController struct {
	service  UserService
	response utility.Response
}

func (c *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	result, status, err := c.service.FindMany()

	if err != nil {
		c.response.InternalServerErrorHandler(w, status, err)
		return
	}

	c.response.GetSuccessResponse(w, nil, result, nil)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
		return
	}

	status, err := c.service.Publish(data, "user.create", "Unauthorize")
	if err != nil {
		c.response.InternalServerErrorHandler(w, status, err)
		return
	}

	c.response.SuccessMessageResponse(w, "Create user with email: "+user.Email)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
		return
	}
	c.service.Publish(data, "user.delete", "Unauthorize")

	userID := strconv.Itoa(int(user.ID))
	c.response.SuccessMessageResponse(w, "Delete user with id: "+userID)
}
