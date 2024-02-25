package user

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/mrspec7er/balky/app/model"
	"github.com/rabbitmq/amqp091-go"
)

type UserListener struct {
	service UserService
}

func (h *UserListener) CreateListener(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
	}

	for data := range messages {
		user := &model.User{}

		err := json.Unmarshal(data.Body, &user)
		if err != nil {
			fmt.Println(err)
		}

		status, err := h.service.CreateService(user)

		if err != nil {
			fmt.Println(status, err)
		}
	}
}
