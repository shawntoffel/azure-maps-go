package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type DailyForecastRequestOptions struct {
	Format          string
	ApiVersion      string
	Details         string
	Duration        *int
	Language        string
	SubscriptionKey string
	Unit            string
	MSClientId      string
}

func (o DailyForecastRequestOptions) Options() Options {
	return Options{
		Format:          o.Format,
		ApiVersion:      o.ApiVersion,
		Details:         o.Details,
		Duration:        o.Duration,
		Language:        o.Language,
		SubscriptionKey: o.SubscriptionKey,
		Unit:            o.Unit,
		MSClientId:      o.MSClientId,
	}
}

type DailyForecastResponse struct {
	Forecasts []entities.DailyForecast      `json:"forecasts,omitempty"`
	Summary   entities.DailyForecastSummary `json:"summary,omitempty"`
}

type DailyIndicesRequestOptions struct {
	Format          string
	ApiVersion      string
	Details         string
	Duration        *int
	IndexGroupId    *int
	IndexId         *int
	Language        string
	SubscriptionKey string
	MSClientId      string
}

func (o DailyIndicesRequestOptions) Options() Options {
	return Options{
		Format:          o.Format,
		ApiVersion:      o.ApiVersion,
		Details:         o.Details,
		Duration:        o.Duration,
		IndexGroupId:    o.IndexGroupId,
		IndexId:         o.IndexId,
		Language:        o.Language,
		SubscriptionKey: o.SubscriptionKey,
		MSClientId:      o.MSClientId,
	}
}

type DailyIndiciesResponse struct {
	Results []entities.DailyIndex `json:"results,omitempty"`
}
