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

type QuarterDayForecastRequestOptions struct {
	Format          string
	ApiVersion      string
	Duration        *int
	Language        string
	SubscriptionKey string
	Unit            string
	MSClientId      string
}

func (o QuarterDayForecastRequestOptions) Options() Options {
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

type QuarterDayForecastResponse struct {
	Forecasts []entities.QuarterDayForecast `json:"forecasts,omitempty"`
}
