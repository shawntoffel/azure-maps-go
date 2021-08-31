package entities

type DailyForecastSummary struct {
	Category  string `json:"category,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	Phrase    string `json:"phrase,omitempty"`
	Severity  int    `json:"severity,omitempty"`
	StartDate string `json:"startDate,omitempty"`
}
