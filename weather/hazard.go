package weather

type Hazards struct {
	HazardDetails  []HazardDetail `json:"hazardDetails,omitempty"`
	MaxHazardIndex int            `json:"maxHazardIndex"`
}
type HazardDetail struct {
	HazardCode  string `json:"hazardCode,omitempty"`
	HazardIndex int    `json:"hazardIndex"`
	ShortPhrase string `json:"shortPhrase,omitempty"`
}
