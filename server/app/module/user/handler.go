package user

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/mrspec7er/balky/app/model"
	"github.com/rabbitmq/amqp091-go"
)

type UserHandler struct {
	service UserService
}

func (h UserHandler) CreateHandler(queue *amqp091.Channel, wg *sync.WaitGroup)  {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, "user.create", "user", true, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
	}

	for data := range messages {
		user := &model.User{}

		err := json.Unmarshal(data.Body, &user)
		if err != nil {
			fmt.Println(err)
		}
		
		h.service.CreateService(user)
	}
}