package master

import (
	"context"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/utility"
)

type MasterReportService struct {
	master    model.MasterReport
	publish   utility.Publisher
	attribute model.Attribute
}

func (s MasterReportService) Create(req *model.MasterReport) (int, error) {
	s.master = *req
	err := s.master.Create()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s MasterReportService) FindMany() ([]model.MasterReport, int, error) {
	masters, err := s.master.FindMany()
	if err != nil {
		return nil, 500, err
	}

	return masters, 201, nil
}

func (s MasterReportService) Delete(req *model.MasterReport) (int, error) {
	s.master = *req
	err := s.master.Delete()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s MasterReportService) Publish(data []byte, queueName string, userEmail string) (int, error) {
	ctx := context.Background()

	payload := utility.Payload{
		Body:      data,
		UserEmail: userEmail,
	}
	err := s.publish.SendMessage(ctx, queueName, &payload)
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s MasterReportService) CreateAttribute(req *model.Attribute) (int, error) {
	s.attribute = *req
	err := s.attribute.Create()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s MasterReportService) FindManyAttribute() ([]model.Attribute, int, error) {
	attributes, err := s.attribute.FindMany()
	if err != nil {
		return nil, 500, err
	}

	return attributes, 201, nil
}

func (s MasterReportService) DeleteAttribute(req *model.Attribute) (int, error) {
	s.attribute = *req
	err := s.attribute.Delete()
	if err != nil {
		return 500, err
	}

	return 201, nil
}
