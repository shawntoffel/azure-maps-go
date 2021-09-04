package weather

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
