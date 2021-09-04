package entities

type WindSpeed struct {
	Unit     string  `json:"unit,omitempty"`
	UnitType int     `json:"unitType"`
	Value    float64 `json:"value"`
}
