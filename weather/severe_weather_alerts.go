package weather

import (
	"net/url"
)

type SevereWeatherAlertsRequestOptions struct {
	Details  string `json:"details,omitempty"`
	Language string `json:"language,omitempty"`
}

func (ccr SevereWeatherAlertsRequestOptions) Encode() string {
	q := url.Values{}

	if ccr.Details != "" {
		q.Add("details", ccr.Details)
	}

	if ccr.Language != "" {
		q.Add("language", ccr.Language)
	}

	return q.Encode()
}
