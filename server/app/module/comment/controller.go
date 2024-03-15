package comment

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/utility"
)

type CommentController struct {
	service  CommentService
	response utility.Response
}

func (c *CommentController) FindAll(w http.ResponseWriter, r *http.Request) {
	appId := chi.URLParam(r, "appId")
	result, status, err := c.service.FindMany(appId)

	if err != nil {
		c.response.InternalServerErrorHandler(w, status, err)
		return
	}

	c.response.GetSuccessResponse(w, nil, result, nil)
}

func (c *CommentController) Create(w http.ResponseWriter, r *http.Request) {
	user, ok := (r.Context().Value(utility.UserContextKey)).(*model.User)
	if !ok {
		c.response.BadRequestHandler(w)
		return
	}

	comment := &model.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	data, err := json.Marshal(comment)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
		return
	}

	status, err := c.service.Publish(data, "comment.create", user.Email)
	if err != nil {
		c.response.InternalServerErrorHandler(w, status, err)
		return
	}

	c.response.SuccessMessageResponse(w, "Create comment")
}

func (c *CommentController) Delete(w http.ResponseWriter, r *http.Request) {
	user, ok := (r.Context().Value(utility.UserContextKey)).(*model.User)
	if !ok {
		c.response.BadRequestHandler(w)
		return
	}
	comment := &model.Comment{}

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	data, err := json.Marshal(comment)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
		return
	}
	c.service.Publish(data, "comment.delete", user.Email)

	c.response.SuccessMessageResponse(w, "Delete comment")
}
