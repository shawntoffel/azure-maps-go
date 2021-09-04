package weather

import "github.com/shawntoffel/azure-maps-go/weather/entities"

type ODataErrorResponse struct {
	Error entities.ODataError `json:"error,omitempty"`
}
