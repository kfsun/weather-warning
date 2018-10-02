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
const fiveMinutes = 300
const twoMinutes = 120

type HKOInfo struct {
	Title       string `xml:"channel>item>title"`
	Description string `xml:"channel>item>description"`
	PubDate     string `xml:"channel>item>pubDate"`
	Guid        string `xml:"channel>item>guid"`
}

func (p *HKOInfo) IsWarningMessage() bool {
	return !strings.Contains(p.Guid, "nowarning")
}

func (p *HKOInfo) GetPublishTimestamp() time.Time {
	t, _ := time.Parse(time.RFC1123Z, p.PubDate)
	return t
}

func (p *HKOInfo) IsNewWarning(previousWarning *warning.WeatherWarning) bool {
	if !previousWarning.PubDate.IsZero() {
		pub := p.GetPublishTimestamp()
		diff := pub.Sub(previousWarning.PubDate)
		if diff.Nanoseconds() == 0 {
			return false
		}
	}

	return true
}

func parse(b []byte) HKOInfo {
	info := HKOInfo{}
	err := xml.Unmarshal(b, &info)
	if err != nil {
		log.Fatalln(err)
	}

	return info
}

func Fetch(c chan warning.WeatherWarning) {
	var oldWarning warning.WeatherWarning

	for {
		xmlBytes := helper.GetHttpContent(rssUrl)
		if xmlBytes == nil {
			// no internet connection, sleep longer and try again
			time.Sleep(time.Second * fiveMinutes)
			continue
		}

		info := parse(xmlBytes)
		if !info.IsWarningMessage() {
			continue
		}

		if info.IsNewWarning(&oldWarning) {
			result := warning.WeatherWarning{
				Title:       info.Title,
				Description: info.Description,
				PubDate:     info.GetPublishTimestamp(),
			}
			oldWarning = result
			c <- result
		}

		time.Sleep(time.Second * twoMinutes)
	}
}
