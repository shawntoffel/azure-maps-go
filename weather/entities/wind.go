package entities

type Wind struct {
	Direction *WindDirection `json:"direction,omitempty"`
	Speed     *WindSpeed     `json:"speed,omitempty"`
}

type WindDirection struct {
	Degrees              float64 `json:"degrees"`
	LocalizedDescription string  `json:"localizedDescription,omitempty"`
}

type WindSpeed struct {
	Unit     string  `json:"unit,omitempty"`
	UnitType int     `json:"unitType"`
	Value    float64 `json:"value"`
}
