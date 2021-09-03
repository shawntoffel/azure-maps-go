package weather

import (
	"net/url"
)

type WeatherAlongRouteRequestOptions struct {
	Language string `json:"language,omitempty"`
}

func (ccr WeatherAlongRouteRequestOptions) Encode() string {
	q := url.Values{}

	if ccr.Language != "" {
		q.Add("language", ccr.Language)
	}

	return q.Encode()
}
