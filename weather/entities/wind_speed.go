package entities

type WindSpeed struct {
	Unit     string  `json:"unit,omitempty"`
	UnitType int     `json:"unitType,omitempty"`
	Value    float64 `json:"value,omitempty"`
}
