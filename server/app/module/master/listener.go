package master

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/mrspec7er/balky/app/model"
	"github.com/rabbitmq/amqp091-go"
)

type MasterReportListener struct {
	service MasterReportService
}

func (l *MasterReportListener) Create(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
	}

	for data := range messages {
		master := &model.MasterReport{}

		err := json.Unmarshal(data.Body, &master)
		if err != nil {
			fmt.Println(err)
		}

		status, err := l.service.Create(master)
		if err != nil {
			fmt.Println(status, err)
		}
	}
}

func (l *MasterReportListener) DeleteListener(queue *amqp091.Channel, wg *sync.WaitGroup, queueName string, consumerTag string) {
	defer wg.Done()

	ctx := context.Background()
	messages, err := queue.ConsumeWithContext(ctx, queueName, consumerTag, true, false, false, false, nil)

	if err != nil {
		fmt.Println(err)
	}

	for data := range messages {
		master := &model.MasterReport{}

		err := json.Unmarshal(data.Body, &master)
		if err != nil {
			fmt.Println(err)
		}

		status, err := l.service.Delete(master)
		if err != nil {
			fmt.Println(status, err)
		}
	}
}
