package config

import (
	"os"

	"github.com/streadway/amqp"
)

func SetupRbMQChannel(rabbitMQInstance *amqp.Connection) error {
	rabbitMqChannelName := os.Getenv("RABBIT_MQ_CHANNEL_NAME")
	rabbitMQChannel, err := rabbitMQInstance.Channel()
	if err != nil {
		return err
	}
	defer rabbitMQChannel.Close()

	_, err = rabbitMQChannel.QueueDeclare(
		rabbitMqChannelName, // queue name
		true,                // durable
		false,               // auto delete
		false,               // exclusive
		false,               // no wait
		nil,                 // arguments
	)
	return err
}
