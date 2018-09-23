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
	Fetch()
}

func (p *WeatherWarning) IsNoWarning() bool {
	return p.PubDate.IsZero()
}
