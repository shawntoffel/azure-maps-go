package entities

type Interval struct {
	CloudCover        int     `json:"cloudCover,omitempty"`
	Color             Color   `json:"color,omitempty"`
	Dbz               float64 `json:"dbz,omitempty"`
	IconCode          int     `json:"iconCode,omitempty"`
	Minute            int     `json:"minute,omitempty"`
	PrecipitationType string  `json:"precipitationType,omitempty"`
	ShortPhrase       string  `json:"shortPhrase,omitempty"`
	SimplifiedColor   Color   `json:"simplifiedColor,omitempty"`
	StartTime         string  `json:"startTime,omitempty"`
	Threshold         string  `json:"threshold,omitempty"`
}
