package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/geekykant/aslipaydekho/scraper/api"
	"github.com/geekykant/aslipaydekho/scraper/utils"
	"github.com/go-co-op/gocron"
)

func main() {
	//Checks required env variables
	if !utils.CheckAllEnvVarsPresent() {
		panic("Exiting. Required env values missing")
	}

	// //Loading .env variables
	// if err := godotenv.Load(); err != nil {
	// 	panic("Error loading .env file")
	// }

	//Defining flags for command line
	var task string
	flag.StringVar(&task, "task", "", "Options - fullpopulate, runcrons")
	flag.Parse()

	switch task {
	case "fullpopulate":
		api.InitPopulateAllCompensationsToMQ()
	case "runcrons":
		runCronJobs()
	default:
		{
			fmt.Println("Exiting. Non-matching flag provided.")
			fmt.Println("Command: go run <filename>.go -task <taskname>")
			flag.PrintDefaults()
			return
		}
	}
}

func runCronJobs() {
	scheduler := gocron.NewScheduler(time.Local)
	fmt.Println(time.Now())
	fmt.Println("[*] Weekly cron scheduled for - 05:00 every sunday.")
	scheduler.Every(1).Sunday().At("05:00").Do(runWeeklyCronTask)
	scheduler.StartBlocking()
}

func runWeeklyCronTask() {
	rabbitMQInstance, err := api.GetRabbitMQInstance()
	if err != nil {
		panic(err)
	}
	defer rabbitMQInstance.Close()

	fmt.Println("[*] Starting weekly cron task at - " + time.Now().Format(time.ANSIC))
	api.FetchPopulateCompensationsSinceLastWeek()
	fmt.Println("[*] Done cron task.")
}
