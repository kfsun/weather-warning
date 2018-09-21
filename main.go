package main

import (
	"github.com/mqu/go-notify"
	"log"
	"os"
	//"time"
	hkprocessor "github.com/kfsworks/weather-warning/processor/hongkong"
)

func main() {
	title, desc, issueTime := hkprocessor.Process()

	if issueTime == nil {
		os.Exit(0)
	}

	log.Println(issueTime)

	notify.Init("Weather Warning")
	warning := notify.NotificationNew(title, desc, "dialog-information")
	warning.Show()
}
