package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

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
	Summary   entities.WeatherAlongRouteSummary `json:"summary,omitempty"`
	Waypoints []entities.Waypoint               `json:"waypoints,omitempty"`
}
