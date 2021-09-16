package weather

type AlertArea struct {
	AlertDetails             string        `json:"alertDetails,omitempty"`
	AlertDetailsLanguageCode string        `json:"alertDetailsLanguageCode,omitempty"`
	EndTime                  string        `json:"endTime,omitempty"`
	LatestStatus             *LatestStatus `json:"latestStatus,omitempty"`
	Name                     string        `json:"name,omitempty"`
	StartTime                string        `json:"startTime,omitempty"`
	Summary                  string        `json:"summary,omitempty"`
}

type SevereWeatherAlert struct {
	AlertAreas  []AlertArea  `json:"alertAreas,omitempty"`
	AlertId     int          `json:"alertId"`
	Category    string       `json:"category,omitempty"`
	Class       string       `json:"class,omitempty"`
	CountryCode string       `json:"countryCode,omitempty"`
	Description *Description `json:"description,omitempty"`
	Disclaimer  string       `json:"disclaimer,omitempty"`
	Level       string       `json:"level,omitempty"`
	Priority    int          `json:"priority"`
	Source      string       `json:"source,omitempty"`
	SourceId    int          `json:"sourceId"`
}
