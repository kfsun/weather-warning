package processor

import (
	"time"
)

type WeatherWarning struct {
	Title       string
	Description string
	PubDate     time.Time
}

type operation interface {
	Process() WeatherWarning
}

func (p *WeatherWarning) IsNoWarning() bool {
	return p.PubDate.IsZero()
}
