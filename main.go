package main

import (
	"fmt"
	//"github.com/kfsworks/weather-warning/processor"
	"github.com/kfsworks/weather-warning/fetcher"
	"github.com/kfsworks/weather-warning/warning"
	"github.com/mqu/go-notify"
	"log"
	"time"
	//    "os"
)

func sendNotification(c chan warning.WeatherWarning) {
	var oldWarning warning.WeatherWarning

	for {
		warning := <-c

		if warning.IsNoWarning() {
			time.Sleep(time.Second * 1)
			continue
		}

		if oldWarning.PubDate.IsZero() {
			log.Println("save old")
			oldWarning = warning
		} else {
			difference := warning.PubDate.Sub(oldWarning.PubDate)
			if difference.Nanoseconds() == 0 {
				log.Println("same and skip")
				continue
			}
		}

		notify.Init("Weather Warning")
		notification := notify.NotificationNew(warning.Title, warning.Description, "dialog-information")
		notification.Show()

		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan warning.WeatherWarning = make(chan warning.WeatherWarning, 5)

	go sendNotification(c)
	go fetcher.Fetch(c)
	//processor.Process()

	var input string
	fmt.Scanln(&input)
}
