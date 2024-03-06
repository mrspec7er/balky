package application

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/module/logger"
	"github.com/rabbitmq/amqp091-go"
)

type ApplicationListener struct {
	service ApplicationService
	logger  logger.LoggerService
}

func (l *ApplicationListener) Create(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		application := &model.Application{}

		userEmail, ok := data.Headers["userEmail"].(string)
		if !ok {
			l.logger.Publish(userEmail, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &application)
		if err != nil {
			l.logger.Publish(userEmail, 400, err.Error())
			continue
		}

		status, err := l.service.Create(application)
		if err != nil {
			l.logger.Publish(userEmail, status, err.Error())
			continue
		}
	}
}

func (l *ApplicationListener) Delete(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		application := &model.Application{}

		userEmail, ok := data.Headers["userEmail"].(string)
		if !ok {
			l.logger.Publish(userEmail, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &application)
		if err != nil {
			l.logger.Publish(userEmail, 400, err.Error())
			continue
		}

		status, err := l.service.Delete(application)
		if err != nil {
			l.logger.Publish(userEmail, status, err.Error())
			continue
		}
	}
}

func (l *ApplicationListener) CreateContent(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		contents := []*model.Content{}

		userEmail, ok := data.Headers["userEmail"].(string)
		if !ok {
			l.logger.Publish(userEmail, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &contents)
		if err != nil {
			l.logger.Publish(userEmail, 400, err.Error())
			continue
		}

		status, err := l.service.CreateContent(contents)
		if err != nil {
			l.logger.Publish(userEmail, status, err.Error())
			continue
		}
	}
}

func (l *ApplicationListener) DeleteContent(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		content := &model.Content{}

		userEmail, ok := data.Headers["userEmail"].(string)
		if !ok {
			l.logger.Publish(userEmail, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &content)
		if err != nil {
			l.logger.Publish(userEmail, 400, err.Error())
			continue
		}

		status, err := l.service.DeleteContent(content)
		if err != nil {
			l.logger.Publish(userEmail, status, err.Error())
			continue
		}
	}
}
