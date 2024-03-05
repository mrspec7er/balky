package user

import (
	"context"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/utility"
)

type UserService struct {
	user    model.User
	publish utility.Publisher
}

func (s UserService) Create(req *model.User) (int, error) {
	s.user = *req
	err := s.user.Create()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s UserService) FindMany() ([]model.User, int, error) {
	users, err := s.user.FindMany()
	if err != nil {
		return nil, 500, err
	}

	return users, 201, nil
}

func (s UserService) FindOne(email string) (*model.User, int, error) {
	s.user.Email = email
	user, err := s.user.FindOne()
	if err != nil {
		return nil, 500, err
	}

	return user, 201, nil
}

func (s UserService) Delete(req *model.User) (int, error) {
	s.user = *req
	err := s.user.Delete()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s UserService) Publish(data []byte, queueName string, userEmail string) (int, error) {
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
