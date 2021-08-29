package entities

type TemeratureSummary struct {
	Past12Hours PastHours `json:"past12Hours,omitempty"`
	Past24Hours PastHours `json:"past24Hours,omitempty"`
	Past6Hours  PastHours `json:"past6Hours,omitempty"`
}
