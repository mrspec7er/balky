package user

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/module/logger"
	"github.com/rabbitmq/amqp091-go"
)

type UserListener struct {
	service UserService
	logger  logger.LoggerService
}

func (l *UserListener) Create(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		user := &model.User{}

		err := json.Unmarshal(data.Body, &user)
		if err != nil {
			l.logger.Publish("Unauthorize", 400, err.Error())
			continue
		}

		status, err := l.service.Create(user)
		if err != nil {
			l.logger.Publish(user.Email, status, err.Error())
			continue
		}
	}
}

func (l *UserListener) DeleteListener(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		user := &model.User{}

		err := json.Unmarshal(data.Body, &user)
		if err != nil {
			l.logger.Publish(user.Email, 400, err.Error())
			continue
		}

		status, err := l.service.Delete(user)
		userEmail, ok := data.Headers["userEmail"].(string)
		if err != nil || !ok {
			l.logger.Publish(userEmail, status, "Missing user credentials")
			continue
		}
	}
}
