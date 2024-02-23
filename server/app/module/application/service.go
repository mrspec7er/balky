package application

import (
	"fmt"
	"os"

	"github.com/mrspec7er/balky/app/model"
	"github.com/rabbitmq/amqp091-go"
)

type ApplicationService struct {
	app model.Application
}

func (s ApplicationService) ListenerService(msgBrokerURI string) (*amqp091.Channel, error) {
	conn, err := amqp091.Dial(os.Getenv("MESSAGE_BROKER_URI"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queue, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return queue, nil
}

func (s ApplicationService) CreateService(req *model.Application)  {
	fmt.Println("RESULT:", req)
}

