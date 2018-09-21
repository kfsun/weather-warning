package main

import (
	"context"
	"encoding/xml"
	"github.com/mqu/go-notify"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

}

type WeatherWarning struct {
	Title       string `xml:"channel>item>title"`
	Description string `xml:"channel>item>description"`
	Guid        string `xml:"channel>item>guid"`
}

func main() {
	log.Println("")

	urltext := "https://rss.weather.gov.hk/rss/WeatherWarningBulletin.xml"

	req, err := http.NewRequest("GET", urltext, nil)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Println(string(body))

	v := WeatherWarning{}
	err = xml.Unmarshal(body, &v)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("result ...")
	log.Println(v)
	//nowarning
	log.Println(" ... done result ")

	notify.Init("Hello world")
	hello := notify.NotificationNew("Hello World!", v.Description, "dialog-information")
	hello.Show()
}
