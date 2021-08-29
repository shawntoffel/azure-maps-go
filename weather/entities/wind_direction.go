package entities

type WindDirection struct {
	Degrees              float64 `json:"degrees,omitempty"`
	LocalizedDescription string  `json:"localizedDescription,omitempty"`
}
