package entities

type HazardDetail struct {
	HazardCode  string `json:"hazardCode,omitempty"`
	HazardIndex int    `json:"hazardIndex"`
	ShortPhrase string `json:"shortPhrase,omitempty"`
}
