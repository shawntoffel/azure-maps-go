package weather

import (
	"net/url"
	"strconv"
)

const DefaultApiVersion = "1.0"
const DefaultFormat = "json"

type Optionable interface {
	Options() Options
}

type Options struct {
	Format          string
	ApiVersion      string
	Details         string
	Duration        *int
	Language        string
	Unit            string
	IndexGroupId    *int
	IndexId         *int
	Interval        *int
	SubscriptionKey string
	MSClientId      string
}

func InitializeOptions(op Optionable, subscriptionKey string, msClientId string) Options {
	options := Options{}
	if op != nil {
		options = op.Options()
	}

	if options.Format == "" {
		options.Format = DefaultFormat
	}

	if options.ApiVersion == "" {
		options.ApiVersion = DefaultApiVersion
	}

	if options.SubscriptionKey == "" {
		options.SubscriptionKey = subscriptionKey
	}

	if options.MSClientId == "" {
		options.MSClientId = msClientId
	}

	return options
}

func (p Options) Encode(query string) string {
	q := url.Values{}

	if p.ApiVersion != "" {
		q.Add("api-version", p.ApiVersion)
	}

	if query != "" {
		q.Add("query", query)
	}

	if p.Details != "" {
		q.Add("details", p.Details)
	}

	if p.Duration != nil {
		q.Add("duration", strconv.Itoa(*p.Duration))
	}

	if p.Language != "" {
		q.Add("language", p.Language)
	}

	if p.Unit != "" {
		q.Add("unit", p.Unit)
	}

	if p.IndexGroupId != nil {
		q.Add("indexGroupId", strconv.Itoa(*p.IndexGroupId))
	}

	if p.IndexId != nil {
		q.Add("indexId", strconv.Itoa(*p.IndexId))
	}

	if p.Interval != nil {
		q.Add("interval", strconv.Itoa(*p.Interval))
	}

	if p.SubscriptionKey != "" {
		q.Add("subscription-key", p.SubscriptionKey)
	}

	return q.Encode()
}
