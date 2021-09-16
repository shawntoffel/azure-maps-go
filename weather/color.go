package weather

type Color struct {
	Blue  int    `json:"blue"`
	Green int    `json:"green"`
	Red   int    `json:"red"`
	Hex   string `json:"hex,omitempty"`
}
