package main

import (
	"time"

	"github.com/geekykant/aslipaydekho/scraper/api"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

func main() {
	//loading env variables
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	// Create a RabbitMQ connection.
	rabbitMQInstance, err := api.GetRabbitMQInstance()
	if err != nil {
		panic(err)
	}
	defer rabbitMQInstance.Close()

	//Intital full populate to MQ
	api.InitPopulateAllCompensationsToMQ()

	//Keep the weekly cron ON in production
	// runCronJobs()
}

func startWeeklyCompensationFetchCronTask() {
	//Calculate the last week time
	//and only fetch those posts in the range

	// api.StartFetchInsertNewCompensationPost()
}

func runCronJobs() {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(5).Seconds().Do(func() {
		startWeeklyCompensationFetchCronTask()
	})
	scheduler.StartBlocking()
}
