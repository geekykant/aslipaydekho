package fetchapi

import (
	"os"
	"sync"

	"github.com/geekykant/aslipaydekho/scraper/config"
	"github.com/streadway/amqp"
)

var rMQLock = &sync.Mutex{}
var rabbitMQInstance *amqp.Connection

func GetRabbitMQInstance() (*amqp.Connection, error) {
	var err error
	if rabbitMQInstance == nil {
		rMQLock.Lock()
		defer rMQLock.Unlock()

		if rabbitMQInstance == nil {
			rabbitMqServerUrl := os.Getenv("RABBIT_MQ_SERVER_URL")
			rabbitMQInstance, err = amqp.Dial(rabbitMqServerUrl)
			if err != nil {
				return nil, err
			}

			//Declares the Service MQ Channel
			err = config.SetupRbMQChannel(rabbitMQInstance)
		}
	}
	return rabbitMQInstance, err
}
