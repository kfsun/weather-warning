package main

import (
	//	"fmt"
	"github.com/kfsworks/weather-warning/fetcher"
	"github.com/kfsworks/weather-warning/warning"
	"github.com/mqu/go-notify"
	"os"
	"os/signal"
	"syscall"
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

//func cleanup() {
//   fmt.Println("cleanup")
//}

func main() {
	var c chan warning.WeatherWarning = make(chan warning.WeatherWarning, 5)

	go sendNotification(c)
	go fetcher.Fetch(c)

	//var input string
	//fmt.Scanln(&input)

	cs := make(chan os.Signal, 2)
	signal.Notify(cs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-cs
		//cleanup()
		//pprof.StopCPUProfile()
		os.Exit(1)
	}()

	for {
		//fmt.Println("sleeping...")
		time.Sleep(10 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}
}
