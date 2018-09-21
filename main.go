package main

import (
	hkprocessor "github.com/kfsworks/weather-warning/processor/hongkong"
	"github.com/mqu/go-notify"
	"os"
)

func main() {
	result := hkprocessor.Process()

	if result.IsNoWarning() {
		os.Exit(0)
	}

	notify.Init("Weather Warning")
	warning := notify.NotificationNew(result.Title, result.Description, "dialog-information")
	warning.Show()
}
