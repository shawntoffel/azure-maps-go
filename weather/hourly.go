package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type HourlyForecastRequestOptions struct {
	Format          string
	ApiVersion      string
	Duration        *int
	Language        string
	SubscriptionKey string
	Unit            string
	MSClientId      string
}

func (o HourlyForecastRequestOptions) Options() Options {
	return Options{
		Format:          o.Format,
		ApiVersion:      o.ApiVersion,
		Duration:        o.Duration,
		Language:        o.Language,
		SubscriptionKey: o.SubscriptionKey,
		Unit:            o.Unit,
		MSClientId:      o.MSClientId,
	}
}

type HourlyForecastResponse struct {
	Forecasts []entities.HourlyForecast `json:"forecasts,omitempty"`
}
