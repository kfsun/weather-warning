/*
 */
package hongkong

import (
	"encoding/xml"
	"github.com/kfsworks/weather-warning/helper"
	"github.com/kfsworks/weather-warning/processor"
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

func Process() processor.WeatherWarning {
	xmlBytes := helper.GetHttpContent(rssUrl)

	info := HKWeatherWarningInfo{}
	err := xml.Unmarshal(xmlBytes, &info)
	if err != nil {
		log.Fatalln(err)
	}

	result := processor.WeatherWarning{Title: info.Title, Description: info.Description}
	if !strings.Contains(info.Guid, "nowarning") {
		t, _ := time.Parse(time.RFC1123Z, info.PubDate)
		result.PubDate = t
	}

	return result
}
