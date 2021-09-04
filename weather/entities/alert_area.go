package entities

type AlertArea struct {
	AlertDetails             string        `json:"alertDetails,omitempty"`
	AlertDetailsLanguageCode string        `json:"alertDetailsLanguageCode,omitempty"`
	EndTime                  string        `json:"endTime,omitempty"`
	LatestStatus             *LatestStatus `json:"latestStatus,omitempty"`
	Name                     string        `json:"name,omitempty"`
	StartTime                string        `json:"startTime,omitempty"`
	Summary                  string        `json:"summary,omitempty"`
}
