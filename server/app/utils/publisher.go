package utils

import (
	"context"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct{}

func (p Publisher) SendMessage(ctx context.Context, queueName string, data []byte) error {
	con, err := amqp091.Dial(os.Getenv("MESSAGE_BROKER_URI"))
	if err != nil {
		return err
	}
	defer con.Close()

	ch, err := con.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	payload := amqp091.Publishing{
		ContentType: "application/json",
		Body:        data,
	}

	err = ch.PublishWithContext(ctx, os.Getenv("EXCHANGE_NAME"), queueName, false, false, payload)
	if err != nil {
		return err
	}

	return nil
}
