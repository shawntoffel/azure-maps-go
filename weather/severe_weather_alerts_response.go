package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type SevereWeatherAlertsResponse struct {
	Results []entities.SevereWeatherAlert `json:"results,omitempty"`
}
