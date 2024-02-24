package user

import (
	"fmt"

	"github.com/mrspec7er/balky/app/model"
)

type UserService struct {
	user model.User
}

func (s UserService) CreateService(req *model.User)  {
	fmt.Println("RESULT USER:", req)
	fmt.Println(s.user.Email)
}

