package entities

type DegreeDaySummary struct {
	Cooling *WeatherUnit `json:"cooling,omitempty"`
	Heating *WeatherUnit `json:"heating,omitempty"`
}
