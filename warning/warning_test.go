package warning

import (
	"testing"
	"time"
)

func TestIsNoWarning(t *testing.T) {
	var w WeatherWarning

	if !w.IsNoWarning() {
		t.Error("Empty WeatherWarning should got true from IsNoWarning.")
	}

	w.PubDate, _ = time.Parse(time.RFC1123Z, "Fri, 21 Sep 2018 19:00:53 +0800")
	w.Title = "Test Title"
	w.Description = "Test Desc"
	if w.IsNoWarning() {
		t.Error("WeatherWarning should got false from IsNoWarning.")
	}
}
