package entities

type WindDirection struct {
	Degrees              float64 `json:"degrees"`
	LocalizedDescription string  `json:"localizedDescription,omitempty"`
}
