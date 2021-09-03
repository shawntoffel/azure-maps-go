package entities

type WeatherAlongRoutePrecipitation struct {
	Dbz  float64 `json:"dbz,omitempty"`
	Type string  `json:"type,omitempty"`
}
