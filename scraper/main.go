package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/geekykant/aslipaydekho/scraper/fetchapi"
	"github.com/geekykant/aslipaydekho/scraper/model"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"

	"github.com/streadway/amqp"
)

func main() {
	//loading env variables
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// Create a RabbitMQ connection.
	rabbitMQInstance, err := fetchapi.GetRabbitMQInstance()
	if err != nil {
		panic(err)
	}
	defer rabbitMQInstance.Close()

	//Intital full populate to MQ
	initPopulateAllCompensationsToMQ()

	//Keep the weekly cron ON in production
	// runCronJobs()
}

func initPopulateAllCompensationsToMQ() {
	paolList, err := fetchapi.FetchCompensationPostsInRange(50, 10)
	if err != nil {
		panic("Error occoured" + err.Error())
	}

	fmt.Printf("[*] Fetched list of count %d\n", len(paolList))
	// fmt.Println(paolList)

	//Insert both - Parsed, Unparsed into MQ
	for _, paol := range paolList {
		err = sendMessageToMQ(&paol)
		if err != nil {
			fmt.Println("Aiyooo can't send message to MQ" + err.Error())
		} else {
			fmt.Printf("[*] Message ID - %s successfully sent to MQ \n", paol.OriginalPost.ID)
		}
	}
}

func sendMessageToMQ(postAndOfferLetter *model.PostAndOfferLetter) error {
	rabbitMQInstance, err := fetchapi.GetRabbitMQInstance()
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
		ContentType: "application/json",
		Body:        paolPayload,
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

func startCompensationFetchCronTask() {
	fetchapi.StartFetchInsertNewCompensationPost()
}

func runCronJobs() {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(5).Seconds().Do(func() {
		startCompensationFetchCronTask()
	})
	scheduler.StartBlocking()
}
