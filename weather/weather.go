package weather

type WeatherAlongRoutePrecipitation struct {
	Dbz  float64 `json:"dbz"`
	Type string  `json:"type,omitempty"`
}

type WeatherUnitRange struct {
	Maximum *WeatherUnit `json:"maximum,omitempty"`
	Minimum *WeatherUnit `json:"minimum,omitempty"`
}

type WeatherUnit struct {
	Unit     string   `json:"unit,omitempty"`
	UnitType UnitType `json:"unitType"`
	Value    float64  `json:"value"`
}
