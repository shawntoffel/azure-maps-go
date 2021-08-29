package weather

import "net/url"

type CurrentConditionsRequest struct {
	ApiVersion      string `json:"apiVersion,omitempty"`
	Query           string `json:"query,omitempty"`
	Details         string `json:"details,omitempty"`
	Duration        int    `json:"duration,omitempty"`
	Language        string `json:"language,omitempty"`
	SubscriptionKey string `json:"subscriptionKey,omitempty"`
	Unit            string `json:"unit,omitempty"`
}

func (ccr CurrentConditionsRequest) Encode() string {
	q := url.Values{}
	if ccr.ApiVersion != "" {
		q.Add("api-version", ccr.ApiVersion)
	}

	if ccr.Query != "" {
		q.Add("query", ccr.Query)
	}

	if ccr.SubscriptionKey != "" {
		q.Add("subscription-key", ccr.SubscriptionKey)
	}

	if ccr.Unit != "" {
		q.Add("unit", ccr.Unit)
	}

	return q.Encode()
}
