package azweather

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
	IntervalSummaries []IntervalSummary     `json:"intervalSummaries,omitempty"`
	Intervals         []Interval            `json:"intervals,omitempty"`
	Summary           MinuteForecastSummary `json:"summary,omitempty"`
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
	Forecasts []DailyForecast      `json:"forecasts,omitempty"`
	Summary   DailyForecastSummary `json:"summary,omitempty"`
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
	Forecasts []HourlyForecast `json:"forecasts,omitempty"`
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
	Forecasts []QuarterDayForecast `json:"forecasts,omitempty"`
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
	Results []DailyIndex `json:"results,omitempty"`
}

type CurrentConditionsRequestOptions struct {
	Format          string
	ApiVersion      string
	Details         string
	Duration        *int
	Language        string
	SubscriptionKey string
	Unit            string
	MSClientId      string
}

func (o CurrentConditionsRequestOptions) Options() Options {
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

type CurrentConditionsResponse struct {
	Results []CurrentConditions `json:"results,omitempty"`
}

type SevereWeatherAlertsRequestOptions struct {
	Format          string
	ApiVersion      string
	Details         string
	Language        string
	SubscriptionKey string
	MSClientId      string
}

func (o SevereWeatherAlertsRequestOptions) Options() Options {
	return Options{
		Format:          o.Format,
		ApiVersion:      o.ApiVersion,
		Details:         o.Details,
		Language:        o.Language,
		SubscriptionKey: o.SubscriptionKey,
		MSClientId:      o.MSClientId,
	}
}

type SevereWeatherAlertsResponse struct {
	Results []SevereWeatherAlert `json:"results,omitempty"`
}

type WeatherAlongRouteRequestOptions struct {
	Format          string
	ApiVersion      string
	Language        string
	SubscriptionKey string
	MSClientId      string
}

func (o WeatherAlongRouteRequestOptions) Options() Options {
	return Options{
		Format:          o.Format,
		ApiVersion:      o.ApiVersion,
		Language:        o.Language,
		SubscriptionKey: o.SubscriptionKey,
		MSClientId:      o.MSClientId,
	}
}

type WeatherAlongRouteResponse struct {
	Summary   WeatherAlongRouteSummary `json:"summary,omitempty"`
	Waypoints []Waypoint               `json:"waypoints,omitempty"`
}
