package entities

type PastHours struct {
	Maximum WeatherUnit `json:"maximum,omitempty"`
	Minimum WeatherUnit `json:"minimum,omitempty"`
}
