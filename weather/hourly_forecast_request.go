package weather

import (
	"net/url"
	"strconv"
)

type HourlyForecastRequestOptions struct {
	Duration *int   `json:"duration,omitempty"`
	Language string `json:"language,omitempty"`
	Unit     string `json:"unit,omitempty"`
}

func (ccr HourlyForecastRequestOptions) Encode() string {
	q := url.Values{}

	if ccr.Duration != nil {
		q.Add("duration", strconv.Itoa(*ccr.Duration))
	}

	if ccr.Language != "" {
		q.Add("language", ccr.Language)
	}

	if ccr.Unit != "" {
		q.Add("unit", ccr.Unit)
	}

	return q.Encode()
}
