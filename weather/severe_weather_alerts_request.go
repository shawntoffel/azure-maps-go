package weather

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
