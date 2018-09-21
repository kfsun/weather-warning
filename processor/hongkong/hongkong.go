/*
 */
package hongkong

import (
	"encoding/xml"
	"github.com/kfsworks/weather-warning/helper"
	"log"
	//"strings"
	"time"
)

const rssUrl = "https://rss.weather.gov.hk/rss/WeatherWarningBulletin.xml"

type HKWeatherWarning struct {
	Title       string `xml:"channel>item>title"`
	Description string `xml:"channel>item>description"`
	PubDate     string `xml:"channel>item>pubDate"`
	Guid        string `xml:"channel>item>guid"`
}

func Process() (string, string, time.Time) {
	xmlBytes := helper.GetHttpContent(rssUrl)

	warning := HKWeatherWarning{}
	err := xml.Unmarshal(xmlBytes, &warning)
	if err != nil {
		log.Fatalln(err)
	}

	t, _ := time.Parse(time.RFC1123Z, warning.PubDate)
	//log.Println(t)

	if strings.Contains(v.Guid, "nowarning") {
		log.Println("no warning")
		return nil, nil, nil
	}

	return warning.Title, warning.Description, t
}
