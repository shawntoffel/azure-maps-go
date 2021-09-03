package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type QuarterDayForecastResponse struct {
	Forecasts []entities.QuarterDayForecast `json:"forecasts,omitempty"`
}
