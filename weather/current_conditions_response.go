package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type CurrentConditionsResponse struct {
	Results []entities.CurrentConditions `json:"results,omitempty"`
}
