package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type MinuteForecastResponse struct {
	IntervalSummaries []entities.IntervalSummary     `json:"intervalSummaries,omitempty"`
	Intervals         []entities.Interval            `json:"intervals,omitempty"`
	Summary           entities.MinuteForecastSummary `json:"summary,omitempty"`
}
