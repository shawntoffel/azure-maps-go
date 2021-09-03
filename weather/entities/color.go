package entities

type Color struct {
	Blue  int    `json:"blue,omitempty"`
	Green int    `json:"green,omitempty"`
	Red   int    `json:"red,omitempty"`
	Hex   string `json:"hex,omitempty"`
}
