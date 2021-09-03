package entities

type SevereWeatherAlert struct {
	AlertAreas  []AlertArea `json:"alertAreas,omitempty"`
	AlertId     int         `json:"alertId,omitempty"`
	Category    string      `json:"category,omitempty"`
	Class       string      `json:"class,omitempty"`
	CountryCode string      `json:"countryCode,omitempty"`
	Description Description `json:"description,omitempty"`
	Disclaimer  string      `json:"disclaimer,omitempty"`
	Level       string      `json:"level,omitempty"`
	Priority    int         `json:"priority,omitempty"`
	Source      string      `json:"source,omitempty"`
	SourceId    int         `json:"sourceId,omitempty"`
}
