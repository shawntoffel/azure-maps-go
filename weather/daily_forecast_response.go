package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type DailyForecastResponse struct {
	Forecasts []entities.DailyForecast      `json:"forecasts,omitempty"`
	Summary   entities.DailyForecastSummary `json:"summary,omitempty"`
}
