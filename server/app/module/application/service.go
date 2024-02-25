package application

import (
	"fmt"

	"github.com/mrspec7er/balky/app/model"
)

type ApplicationService struct {
	app model.Application
}

func (s ApplicationService) Create(req *model.Application) {
	fmt.Println("RESULT APP:", req)
	fmt.Println(s.app.Number)
}
