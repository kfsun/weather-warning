package processor

import (
	"github.com/kfsworks/weather-warning/fetcher"
	"github.com/kfsworks/weather-warning/warning"
	"github.com/mqu/go-notify"
	"time"
	//    "os"
)

/*
type operation interface {
    Fetch()
}
*/
//var fetcher map[string]string
//func init() {
//	fetcher = make(map[string]string)
//	fetcher["HKO"]
//}

func sendNotification(c chan warning.WeatherWarning) {
	for {
		warning := <-c

		if warning.IsNoWarning() {
			time.Sleep(time.Second * 1)
			continue
		}

		notify.Init("Weather Warning")
		notification := notify.NotificationNew(warning.Title, warning.Description, "dialog-information")
		notification.Show()

		time.Sleep(time.Second * 1)
	}
}

func Process() {
	var c chan warning.WeatherWarning = make(chan warning.WeatherWarning, 5)

	go sendNotification(c)
	go fetcher.Fetch(c)
}
