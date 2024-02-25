package user

import (
	"github.com/mrspec7er/balky/app/model"
)

type UserService struct {
	user model.User
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

func (s UserService) Delete(req *model.User) (int, error) {
	s.user = *req
	err := s.user.Delete()
	if err != nil {
		return 500, err
	}

	return 201, nil
}
