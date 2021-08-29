package entities

type Wind struct {
	Direction WindDirection `json:"direction,omitempty"`
	Speed     WindSpeed     `json:"speed,omitempty"`
}
