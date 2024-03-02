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

		userId, ok := data.Headers["userId"].(string)
		if !ok {
			l.logger.Publish(userId, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &master)
		if err != nil {
			l.logger.Publish(userId, 400, err.Error())
			continue
		}

		status, err := l.service.Create(master)
		if err != nil {
			l.logger.Publish(userId, status, err.Error())
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

		userId, ok := data.Headers["userId"].(string)
		if !ok {
			l.logger.Publish(userId, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &master)
		if err != nil {
			l.logger.Publish(userId, 400, err.Error())
			continue
		}

		status, err := l.service.Delete(master)
		if err != nil {
			l.logger.Publish(userId, status, err.Error())
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
		attribute := &model.Attribute{}

		userId, ok := data.Headers["userId"].(string)
		if !ok {
			l.logger.Publish(userId, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &attribute)
		if err != nil {
			l.logger.Publish(userId, 400, err.Error())
			continue
		}

		status, err := l.service.CreateAttribute(attribute)
		if err != nil {
			l.logger.Publish(userId, status, err.Error())
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

		userId, ok := data.Headers["userId"].(string)
		if !ok {
			l.logger.Publish(userId, 400, "Missing user credentials")
			continue
		}

		err := json.Unmarshal(data.Body, &attribute)
		if err != nil {
			l.logger.Publish(userId, 400, err.Error())
			continue
		}

		status, err := l.service.DeleteAttribute(attribute)
		if err != nil {
			l.logger.Publish(userId, status, err.Error())
			continue
		}
	}
}
