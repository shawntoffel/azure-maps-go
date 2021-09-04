package weather

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
