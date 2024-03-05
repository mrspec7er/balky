package utility

import (
	"context"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

type Payload struct {
	Body      []byte
	UserEmail string
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
			"userEmail": p.UserEmail,
		},
	}

	err = ch.PublishWithContext(ctx, os.Getenv("EXCHANGE_NAME"), queueName, false, false, payload)
	if err != nil {
		return err
	}

	return nil
}

func (Publisher) Log(ctx context.Context, data []byte) error {
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

	err = ch.PublishWithContext(ctx, os.Getenv("LOGGER_EXCHANGE"), os.Getenv("LOGGER_QUEUE"), false, false, payload)
	if err != nil {
		return err
	}

	return nil
}
