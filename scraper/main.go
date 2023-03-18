package main

import (
	"time"

	"github.com/geekykant/aslipaydekho/scraper/config"
	"github.com/geekykant/aslipaydekho/scraper/fetchapi"
	"github.com/go-co-op/gocron"
	"github.com/sethgrid/pester"
)

func main() {
	//loading env variables
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// secretKey := os.Getenv("SECRET_KEY")
	// fmt.Println(secretKey)

	//Intital full populate to MQ
	forceFetchAndPushAllCompensationsToMQ()

	// runCronJobs()
}

func forceFetchAndPushAllCompensationsToMQ() {
	//init http config
	client := pester.New()
	config.SetHttpReqConfig(client)

	//
}

func startCompensationFetchCronTask(client *pester.Client) {
	fetchapi.FetchLeetCodeCompensationPosts(client)
}

func runCronJobs() {
	//init http config
	client := pester.New()
	config.SetHttpReqConfig(client)

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(5).Seconds().Do(func() {
		startCompensationFetchCronTask(client)
	})
	scheduler.StartBlocking()
}
