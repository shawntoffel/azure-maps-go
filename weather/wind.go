package azweather

type Wind struct {
	Direction *WindDirection `json:"direction,omitempty"`
	Speed     *WeatherUnit   `json:"speed,omitempty"`
}

type WindDirection struct {
	Degrees              float64 `json:"degrees"`
	LocalizedDescription string  `json:"localizedDescription,omitempty"`
}
