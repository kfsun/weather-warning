package main

import (
	"encoding/xml"
	"github.com/kfsworks/weather-warning/helper"
	"github.com/mqu/go-notify"
	"log"
	"os"
	"strings"
)

func init() {

	// go get github.com/mqu/go-notify
	//https://rss.weather.gov.hk/
	//https://rss.weather.gov.hk/rss/WeatherWarningBulletin.xml
	//https://rss.weather.gov.hk/rss/WeatherWarningSummaryv2.xml
	//https://rss.weather.gov.hk/rss/LocalWeatherForecast.xml
	//https://rss.weather.gov.hk/rss/CurrentWeather.xml
	//https://rss.weather.gov.hk/rss/SeveralDaysWeatherForecast.xml
	//https://rss.weather.gov.hk/rss/QuickEarthquakeMessage.xml
	//https://rss.weather.gov.hk/rss/FeltEarthquake.xml

	//<pubDate>Fri, 21 Sep 2018 11:45:00 +0800</pubDate>
}

type WeatherWarning struct {
	Title       string `xml:"channel>item>title"`
	Description string `xml:"channel>item>description"`
	PubDate
	Guid string `xml:"channel>item>guid"`
}

func main() {
	urltext := "https://rss.weather.gov.hk/rss/WeatherWarningBulletin.xml"
	body := helper.GetHttpContent(urltext)

	v := WeatherWarning{}
	err := xml.Unmarshal(body, &v)
	if err != nil {
		log.Fatalln(err)
	}

	if strings.Contains(v.Guid, "nowarning") {
		os.Exit(0)
	}

	notify.Init("Weather Warning")
	//warning := notify.NotificationNew("Weather Warning!", v.Description, "dialog-information")
	warning := notify.NotificationNew(v.Title, v.Description, "dialog-information")
	warning.Show()
}
