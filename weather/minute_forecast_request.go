package weather

import (
	"net/url"
	"strconv"
)

type MinuteForecastRequestOptions struct {
	Interval *int   `json:"duration,omitempty"`
	Language string `json:"language,omitempty"`
}

func (ccr MinuteForecastRequestOptions) Encode() string {
	q := url.Values{}

	if ccr.Interval != nil {
		q.Add("interval", strconv.Itoa(*ccr.Interval))
	}

	if ccr.Language != "" {
		q.Add("language", ccr.Language)
	}

	return q.Encode()
}
