package entities

type IntervalSummary struct {
	BriefPhrase  string `json:"briefPhrase,omitempty"`
	EndMinute    int    `json:"endMinute,omitempty"`
	IconCode     int    `json:"iconCode,omitempty"`
	LongPhrase   string `json:"longPhrase,omitempty"`
	ShortPhrase  string `json:"shortPhrase,omitempty"`
	StartMinute  int    `json:"startMinute,omitempty"`
	TotalMinutes int    `json:"totalMinutes,omitempty"`
}
