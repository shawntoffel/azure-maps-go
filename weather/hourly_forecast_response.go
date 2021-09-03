package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type HourlyForecastResponse struct {
	Forecasts []entities.HourlyForecast `json:"forecasts,omitempty"`
}
