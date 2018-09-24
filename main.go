package main

import (
	"fmt"
	"github.com/kfsworks/weather-warning/fetcher"
	"github.com/kfsworks/weather-warning/warning"
	"github.com/mqu/go-notify"
	"time"
)

func sendNotification(c chan warning.WeatherWarning) {
	for {
		warning := <-c

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

	var input string
	fmt.Scanln(&input)
}
