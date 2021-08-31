package weather

import "net/url"

type Encodable interface {
	Encode() string
}

type Request struct {
	ApiVersion      string `json:"apiVersion,omitempty"`
	Query           string `json:"query,omitempty"`
	SubscriptionKey string `json:"subscriptionKey,omitempty"`
}

func (r Request) Encode() string {
	urlValues := url.Values{}

	if r.ApiVersion != "" {
		urlValues.Add("api-version", r.ApiVersion)
	}

	if r.Query != "" {
		urlValues.Add("query", r.Query)
	}

	if r.SubscriptionKey != "" {
		urlValues.Add("subscription-key", r.SubscriptionKey)
	}

	return urlValues.Encode()
}
