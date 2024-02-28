package utils

import (
	"context"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

type Payload struct {
	Body   []byte
	UserID string
}

type Publisher struct{}

func (Publisher) SendMessage(ctx context.Context, queueName string, p *Payload) error {
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
		Body:        p.Body,
		Headers: amqp091.Table{
			"userId": p.UserID,
		},
	}

	err = ch.PublishWithContext(ctx, os.Getenv("EXCHANGE_NAME"), queueName, false, false, payload)
	if err != nil {
		return err
	}

	return nil
}
