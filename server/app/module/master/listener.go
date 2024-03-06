package master

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/module/logger"
	"github.com/rabbitmq/amqp091-go"
)

type MasterReportListener struct {
	service MasterReportService
	logger  logger.LoggerService
}

func (l *MasterReportListener) Create(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		master := &model.MasterReport{}

		userEmail, ok := data.Headers["userEmail"].(string)
		if !ok {
			l.logger.Publish(userEmail, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &master)
		if err != nil {
			l.logger.Publish(userEmail, 400, err.Error())
			continue
		}

		status, err := l.service.Create(master)
		if err != nil {
			l.logger.Publish(userEmail, status, err.Error())
			continue
		}
	}
}

func (l *MasterReportListener) Delete(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		master := &model.MasterReport{}

		userEmail, ok := data.Headers["userEmail"].(string)
		if !ok {
			l.logger.Publish(userEmail, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &master)
		if err != nil {
			l.logger.Publish(userEmail, 400, err.Error())
			continue
		}

		status, err := l.service.Delete(master)
		if err != nil {
			l.logger.Publish(userEmail, status, err.Error())
			continue
		}
	}
}

func (l *MasterReportListener) CreateAttribute(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		attributes := []*model.Attribute{}

		userEmail, ok := data.Headers["userEmail"].(string)
		if !ok {
			l.logger.Publish(userEmail, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &attributes)
		if err != nil {
			l.logger.Publish(userEmail, 400, err.Error())
			continue
		}

		status, err := l.service.CreateAttribute(attributes)
		if err != nil {
			l.logger.Publish(userEmail, status, err.Error())
			continue
		}
	}
}

func (l *MasterReportListener) DeleteAttribute(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		l.logger.Publish("Server", 500, err.Error())
	}

	for data := range messages {
		attribute := &model.Attribute{}

		userEmail, ok := data.Headers["userEmail"].(string)
		if !ok {
			l.logger.Publish(userEmail, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &attribute)
		if err != nil {
			l.logger.Publish(userEmail, 400, err.Error())
			continue
		}

		status, err := l.service.DeleteAttribute(attribute)
		if err != nil {
			l.logger.Publish(userEmail, status, err.Error())
			continue
		}
	}
}
