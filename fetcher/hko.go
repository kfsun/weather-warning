/*
 */
package fetcher

import (
	"encoding/xml"
	"github.com/kfsworks/weather-warning/helper"
	"github.com/kfsworks/weather-warning/warning"
	"log"
	"strings"
	"time"
)

const rssUrl = "https://rss.weather.gov.hk/rss/WeatherWarningBulletin.xml"

type HKWeatherWarningInfo struct {
	Title       string `xml:"channel>item>title"`
	Description string `xml:"channel>item>description"`
	PubDate     string `xml:"channel>item>pubDate"`
	Guid        string `xml:"channel>item>guid"`
}

func parse(b []byte) HKWeatherWarningInfo {
	info := HKWeatherWarningInfo{}
	err := xml.Unmarshal(b, &info)
	if err != nil {
		log.Fatalln(err)
	}

	return info
}

func Fetch(c chan warning.WeatherWarning) {
	for {
		xmlBytes := helper.GetHttpContent(rssUrl)

		info := parse(xmlBytes)

		result := warning.WeatherWarning{Title: info.Title, Description: info.Description}
		if !strings.Contains(info.Guid, "nowarning") {
			t, _ := time.Parse(time.RFC1123Z, info.PubDate)
			result.PubDate = t
		}

		c <- result
		time.Sleep(time.Second * 10)
	}
}
