package weather

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
