package api

import (
	"encoding/json"
	"os"

	"github.com/geekykant/aslipaydekho/scraper/model"
	"github.com/streadway/amqp"
)

func SendOfferLetterToMQ(postAndOfferLetter *model.PostAndOfferLetter) error {
	rabbitMQInstance, err := GetRabbitMQInstance()
	if err != nil {
		return err
	}

	rabbitMQChannel, err := rabbitMQInstance.Channel()
	if err != nil {
		return err
	}
	defer rabbitMQChannel.Close()

	//PostAndOfferLetter parsed to json format
	paolPayload, err := json.Marshal(postAndOfferLetter)
	if err != nil {
		return err
	}

	// Create a message to publish.
	message := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         paolPayload,
	}

	// Attempt to publish a message to the queue.
	rabbitMqChannelName := os.Getenv("RABBIT_MQ_CHANNEL_NAME")
	if err := rabbitMQChannel.Publish(
		"",                  // exchange
		rabbitMqChannelName, // queue name
		false,               // mandatory
		false,               // immediate
		message,             // message to publish
	); err != nil {
		return err
	}
	return nil
}
