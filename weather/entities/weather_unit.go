package entities

type WeatherUnit struct {
	Unit     string   `json:"unit,omitempty"`
	UnitType UnitType `json:"unitType"`
	Value    float64  `json:"value"`
}
