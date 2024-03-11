package application

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/utility"
)

type ApplicationController struct {
	service  ApplicationService
	response utility.Response
}

func (c *ApplicationController) FindAll(w http.ResponseWriter, r *http.Request) {

	result, status, err := c.service.FindMany()

	if err != nil {
		c.response.InternalServerErrorHandler(w, status, err)
		return
	}

	c.response.GetSuccessResponse(w, nil, result, nil)
}

func (c *ApplicationController) Create(w http.ResponseWriter, r *http.Request) {
	user, ok := (r.Context().Value(utility.UserContextKey)).(*model.User)
	if !ok {
		c.response.BadRequestHandler(w)
		return
	}
	app := &model.Application{}

	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	app.UserEmail = user.Email

	data, err := json.Marshal(app)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
		return
	}

	c.service.Publish(data, "app.create", user.Email)

	c.response.SuccessMessageResponse(w, "Create application with number: "+app.Number)

}

func (c *ApplicationController) Delete(w http.ResponseWriter, r *http.Request) {
	app := &model.Application{}

	if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	data, err := json.Marshal(app)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
		return
	}

	user, ok := (r.Context().Value(utility.UserContextKey)).(*model.User)
	if !ok {
		c.response.BadRequestHandler(w)
		return
	}

	c.service.Publish(data, "app.delete", user.Email)

	c.response.SuccessMessageResponse(w, "Delete app with number: "+app.Number)
}

func (c *ApplicationController) FindAllContent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	result, status, err := c.service.FindManyContent(id)

	if err != nil {
		c.response.InternalServerErrorHandler(w, status, err)
		return
	}

	c.response.GetSuccessResponse(w, nil, result, nil)
}

func (c *ApplicationController) CreateContent(w http.ResponseWriter, r *http.Request) {
	contents := []*model.Content{}

	if err := json.NewDecoder(r.Body).Decode(&contents); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	data, err := json.Marshal(contents)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
		return
	}

	user, ok := (r.Context().Value(utility.UserContextKey)).(*model.User)
	if !ok {
		c.response.BadRequestHandler(w)
		return
	}

	c.service.Publish(data, "content.create", user.Email)

	c.response.SuccessMessageResponse(w, "Create content")
}

func (c *ApplicationController) DeleteContent(w http.ResponseWriter, r *http.Request) {
	content := &model.Content{}

	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	data, err := json.Marshal(content)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
		return
	}

	user, ok := (r.Context().Value(utility.UserContextKey)).(*model.User)
	if !ok {
		c.response.BadRequestHandler(w)
		return
	}

	c.service.Publish(data, "content.delete", user.Email)

	c.response.SuccessMessageResponse(w, "Delete content")
}

type InsertReaction struct {
	ApplicationNumber string `json:"applicationNumber"`
	UserEmail         string `json:"userEmail"`
}

func (c *ApplicationController) CreateReaction(w http.ResponseWriter, r *http.Request) {
	user, ok := (r.Context().Value(utility.UserContextKey)).(*model.User)
	if !ok {
		c.response.BadRequestHandler(w)
		return
	}
	reaction := &InsertReaction{}

	if err := json.NewDecoder(r.Body).Decode(&reaction); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	data, err := json.Marshal(reaction)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
		return
	}

	c.service.Publish(data, "reaction.create", user.Email)

	c.response.SuccessMessageResponse(w, "Create reaction: "+reaction.ApplicationNumber)

}
