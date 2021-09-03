package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type DailyIndiciesResponse struct {
	Results []entities.DailyIndex `json:"results,omitempty"`
}
