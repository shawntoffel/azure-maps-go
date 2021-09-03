package entities

type Notification struct {
	HazardCode  string `json:"hazardCode,omitempty"`
	HazardIndex int    `json:"hazardIndex,omitempty"`
	ShortPhrase string `json:"shortPhrase,omitempty"`
	Type        string `json:"type,omitempty"`
}
