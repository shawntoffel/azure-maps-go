package azweather

type Notification struct {
	HazardCode  string `json:"hazardCode,omitempty"`
	HazardIndex int    `json:"hazardIndex"`
	ShortPhrase string `json:"shortPhrase,omitempty"`
	Type        string `json:"type,omitempty"`
}
