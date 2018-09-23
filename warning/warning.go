package warning

import (
	"time"
)

type WeatherWarning struct {
	Title       string
	Description string
	PubDate     time.Time
}

func (p *WeatherWarning) IsNoWarning() bool {
	return p.PubDate.IsZero()
}
