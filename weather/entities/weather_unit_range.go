package entities

type WeatherUnitRange struct {
	Maximum *WeatherUnit `json:"maximum,omitempty"`
	Minimum *WeatherUnit `json:"minimum,omitempty"`
}
