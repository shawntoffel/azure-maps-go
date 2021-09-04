package entities

type IntervalSummary struct {
	BriefPhrase  string `json:"briefPhrase,omitempty"`
	EndMinute    int    `json:"endMinute"`
	IconCode     int    `json:"iconCode"`
	LongPhrase   string `json:"longPhrase,omitempty"`
	ShortPhrase  string `json:"shortPhrase,omitempty"`
	StartMinute  int    `json:"startMinute"`
	TotalMinutes int    `json:"totalMinutes"`
}
