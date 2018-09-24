package fetcher

import (
	"github.com/kfsworks/weather-warning/warning"
	"testing"
	"time"
)

const rssNonWarning = `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet href="styleb.xsl" type="text/xsl"?>
    <rss version="2.0">
        <channel>
            <title>Weather Warning Information</title>
            <description>Weather Warning Information provided by the Hong Kong Observatory</description>
            <link>http://www.weather.gov.hk/textonly/warning/detail.htm</link>
            <language>en-us</language>
            <webMaster>mailbox@hko.gov.hk</webMaster>
            <copyright>The content available in this file, including but not limited to all text, graphics, drawings,diagrams, photographs and compilation of data or other materials are protected by copyright. The Government of the Hong Kong Special Administrative Region is the owner of all copyright works contained in this website.</copyright>
            <pubDate>Fri, 21 Sep 2018 19:00:53 +0800</pubDate>
            <image>
                <title>Hong Kong Observatory</title>
                <url>http://rss.weather.gov.hk/img/logo_dblue.gif</url>
                <width>333</width>
                <height>65</height>
            </image>
            <item>
                <title><![CDATA[There is no special announcement (19:00 HKT on 21.09.2018)]]></title>
                <link>http://www.weather.gov.hk/textonly/warning/detail.htm</link>
                <author>Hong Kong Observatory</author>
                <description><![CDATA[There is no special announcement (19:00 HKT on 21.09.2018)]]></description>
                <pubDate>Fri, 21 Sep 2018 19:00:53 +0800</pubDate>
                <guid isPermaLink="false">http://rss.weather.gov.hk/rss/1537527653/nowarning</guid>
            </item>
        </channel>
    </rss>`

const rssWarning = `<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet href="styleb.xsl" type="text/xsl"?>
    <rss version="2.0">
        <channel>
            <title>Weather Warning Information</title>
            <description>Weather Warning Information provided by the Hong Kong Observatory</description>
            <link>http://www.weather.gov.hk/textonly/warning/detail.htm</link>
            <language>en-us</language>
            <webMaster>mailbox@hko.gov.hk</webMaster>
            <copyright>The content available in this file, including but not limited to all text, graphics, drawings,diagrams, photographs and compilation of data or other materials are protected by copyright. The Government of the Hong Kong Special Administrative Region is the owner of all copyright works contained in this website.</copyright>
            <pubDate>Fri, 21 Sep 2018 11:48:33 +0800</pubDate>
            <image>
                <title>Hong Kong Observatory</title>
                <url>http://rss.weather.gov.hk/img/logo_dblue.gif</url>
                <width>333</width>
                <height>65</height>
            </image>
            <item>
                <title><![CDATA[Information about Very Hot Weather Warning (11:45 HKT on 21.09.2018)]]></title>
                <link>http://www.weather.gov.hk/textonly/warning/detail.htm</link>
                <author>Hong Kong Observatory</author>
                <description><![CDATA[The Very Hot Weather Warning has been issued by the Hong Kong Observatory at 11:45 a.m.The Hong Kong Observatory is forecasting very hot weather in Hong Kong today. The risk of heatstroke is high.<br/><br/>When engaged in outdoor work or activities, drink plenty of water and avoid over exertion. If not feeling well, take a rest in the shade or cooler place as soon as possible.<br/><br/>People staying indoors without air-conditioning should keep windows open as far as possible to ensure that there is adequate ventilation.<br/><br/>Avoid prolonged exposure under sunlight. Loose clothing, suitable hats and UV-blocking sunglasses can reduce the chance of sunburn by solar ultraviolet radiation.<br/><br/>Swimmers and those taking part in outdoor activities should use a sunscreen lotion of SPF 15 or above, and should re-apply it frequently.<br/><br/>Beware of health and wellbeing of elderly or persons with chronic medical conditions. If you know of them, call or visit them occasionally to check if they need any assistance.<br/><br/>Dispatched by Hong Kong Observatory at 11:45 HKT on 21.09.2018<br/><br/>]]></description>
                <pubDate>Fri, 21 Sep 2018 11:45:00 +0800</pubDate>
                <guid isPermaLink="false">http://rss.weather.gov.hk/rss/201809211145Information about Very Hot Weather Warning1041</guid>
            </item>
        </channel>
    </rss>`

func TestNonWarningRss(t *testing.T) {
	info := parse([]byte(rssNonWarning))

	if info.IsWarningMessage() {
		t.Error("This is not a warning message")
	}
}

func TestWarningRss(t *testing.T) {
	info := parse([]byte(rssWarning))

	if !info.IsWarningMessage() {
		t.Error("This is a warning message")
	}

	tt, _ := time.Parse(time.RFC1123Z, "Fri, 21 Sep 2018 11:45:00 +0800")
	pubtime := info.GetPublishTimestamp()
	if pubtime.Sub(tt).Nanoseconds() != 0 {
		t.Error("PubTime is not correct.")
	}

	var oldWarning warning.WeatherWarning
	if !info.IsNewWarning(&oldWarning) {
		t.Error("Case 1: This is a new warning since the oldWarning is empty.")
	}

	oldWarning.Title = info.Title
	oldWarning.Description = info.Description
	oldWarning.PubDate = info.GetPublishTimestamp()
	if info.IsNewWarning(&oldWarning) {
		t.Error("Case 2: This is not a new warning since the oldWarning is updated with the same struct.")
	}
}
