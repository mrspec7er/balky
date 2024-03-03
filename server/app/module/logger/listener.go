package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/mrspec7er/balky/app/model"
	"github.com/rabbitmq/amqp091-go"
)

type HeaderPayload struct {
	UserEmail string `json:"userEmail"`
}

type LoggerListener struct {
	service LoggerService
}

func (l *LoggerListener) Create(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
	}

	for data := range messages {

		log := &model.Logger{}
		err := json.Unmarshal(data.Body, &log)
		if err != nil {
			fmt.Println(err)
		}

		err = l.service.Create(log)
		if err != nil {
			panic(err)
		}
	}
}
