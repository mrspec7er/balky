package comment

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/module/logger"
	"github.com/rabbitmq/amqp091-go"
)

type CommentListener struct {
	service CommentService
	logger  logger.LoggerService
}

func (l *CommentListener) Create(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		userEmail, ok := data.Headers["userEmail"].(string)
		if !ok {
			l.logger.Publish(userEmail, 400, "Missing user credentials")
			continue
		}

		comment := &model.Comment{}

		err := json.Unmarshal(data.Body, &comment)
		if err != nil {
			l.logger.Publish("Unauthorize", 400, err.Error())
			continue
		}

		status, err := l.service.Create(comment)
		if err != nil {
			l.logger.Publish(userEmail, status, err.Error())
			continue
		}
	}
}

func (l *CommentListener) DeleteListener(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		userEmail, ok := data.Headers["userEmail"].(string)
		if !ok {
			l.logger.Publish(userEmail, 400, "Missing user credentials")
			continue
		}

		comment := &model.Comment{}

		err := json.Unmarshal(data.Body, &comment)
		if err != nil {
			l.logger.Publish(userEmail, 400, err.Error())
			continue
		}

		status, err := l.service.Delete(comment)
		if err != nil || !ok {
			l.logger.Publish(userEmail, status, err.Error())
			continue
		}
	}
}
