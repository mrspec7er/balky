package application

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mrspec7er/balky/app/model"
	"github.com/rabbitmq/amqp091-go"
)

type ApplicationHandler struct {
	service ApplicationService
}

func (h ApplicationHandler) CreateHandler()  {
	conn, err := amqp091.Dial(os.Getenv("MESSAGE_BROKER_URI"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queue, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, "create", "application", true, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
	}

	for data := range messages {
		app := &model.Application{}

		err := json.Unmarshal(data.Body, &app)
		if err != nil {
			fmt.Println(err)
		}
		
		h.service.CreateService(app)
	}
}