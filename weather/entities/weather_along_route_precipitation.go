package entities

type WeatherAlongRoutePrecipitation struct {
	Dbz  float64 `json:"dbz"`
	Type string  `json:"type,omitempty"`
}
