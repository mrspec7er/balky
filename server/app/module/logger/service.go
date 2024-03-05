package logger

import (
	"context"
	"encoding/json"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/utility"
)

type LoggerService struct {
	logger  model.Logger
	publish utility.Publisher
}

func (s LoggerService) Create(req *model.Logger) error {
	s.logger = *req
	err := s.logger.Create()
	if err != nil {
		return err
	}

	return nil
}

func (s LoggerService) FindMany() ([]model.Logger, int, error) {
	Loggers, err := s.logger.FindMany()
	if err != nil {
		return nil, 500, err
	}

	return Loggers, 201, nil
}

func (s LoggerService) Delete(req *model.Logger) (int, error) {
	s.logger = *req
	err := s.logger.Delete()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s LoggerService) Publish(userEmail string, status int, message string) {
	ctx := context.Background()

	s.logger.Author = userEmail
	s.logger.Message = message

	data := model.Logger{
		Author:  userEmail,
		Status:  status,
		Message: message,
	}

	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = s.publish.Log(ctx, body)
	if err != nil {
		panic(err)
	}
}
