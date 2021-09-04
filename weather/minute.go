package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type MinuteForecastRequestOptions struct {
	Format          string
	ApiVersion      string
	Interval        *int
	Language        string
	SubscriptionKey string
	MSClientId      string
}

func (o MinuteForecastRequestOptions) Options() Options {
	return Options{
		Format:          o.Format,
		ApiVersion:      o.ApiVersion,
		Interval:        o.Interval,
		Language:        o.Language,
		SubscriptionKey: o.SubscriptionKey,
		MSClientId:      o.MSClientId,
	}
}

type MinuteForecastResponse struct {
	IntervalSummaries []entities.IntervalSummary     `json:"intervalSummaries,omitempty"`
	Intervals         []entities.Interval            `json:"intervals,omitempty"`
	Summary           entities.MinuteForecastSummary `json:"summary,omitempty"`
}
