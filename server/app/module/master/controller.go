package master

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/utils"
)

type MasterReportController struct {
	service  MasterReportService
	response utils.Response
}

func (c *MasterReportController) FindAll(w http.ResponseWriter, r *http.Request) {
	result, status, err := c.service.FindMany()

	if err != nil {
		c.response.InternalServerErrorHandler(w, status, err)
		return
	}

	c.response.GetSuccessResponse(w, nil, result, nil)
}

func (c *MasterReportController) Create(w http.ResponseWriter, r *http.Request) {
	master := &model.MasterReport{}

	if err := json.NewDecoder(r.Body).Decode(&master); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	data, err := json.Marshal(master)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
		return
	}

	userId := "dummy"
	c.service.Publish(data, "master.create", userId)

	c.response.SuccessMessageResponse(w, "Create master with name: "+master.Name)
}

func (c *MasterReportController) Delete(w http.ResponseWriter, r *http.Request) {
	master := &model.MasterReport{}

	if err := json.NewDecoder(r.Body).Decode(&master); err != nil {
		c.response.BadRequestHandler(w)
		return
	}

	data, err := json.Marshal(master)
	if err != nil {
		c.response.InternalServerErrorHandler(w, 500, err)
		return
	}

	userId := "dummy"
	c.service.Publish(data, "master.delete", userId)

	masterID := strconv.Itoa(int(master.ID))
	c.response.SuccessMessageResponse(w, "Delete master with id: "+masterID)
}
