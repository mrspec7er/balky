package application

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/mrspec7er/balky/app/model"
	"github.com/rabbitmq/amqp091-go"
)

type ApplicationHandler struct {
	service ApplicationService
}

func (h ApplicationHandler) CreateHandler(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

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
