package weather

import (
	"net/url"
	"strconv"
)

type DailyForecastRequestOptions struct {
	Details  string `json:"details,omitempty"`
	Duration *int   `json:"duration,omitempty"`
	Language string `json:"language,omitempty"`
	Unit     string `json:"unit,omitempty"`
}

func (ccr DailyForecastRequestOptions) Encode() string {
	q := url.Values{}

	if ccr.Details != "" {
		q.Add("details", ccr.Details)
	}

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
