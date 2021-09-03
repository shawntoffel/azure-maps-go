package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type WeatherAlongRouteResponse struct {
	Summary   entities.WeatherAlongRouteSummary `json:"summary,omitempty"`
	Waypoints []entities.Waypoint               `json:"waypoints,omitempty"`
}
