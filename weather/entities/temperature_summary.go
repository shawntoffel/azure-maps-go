package entities

type TemeratureSummary struct {
	Past12Hours WeatherUnitRange `json:"past12Hours,omitempty"`
	Past24Hours WeatherUnitRange `json:"past24Hours,omitempty"`
	Past6Hours  WeatherUnitRange `json:"past6Hours,omitempty"`
}
