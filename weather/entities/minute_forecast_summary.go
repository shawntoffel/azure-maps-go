package entities

type MinuteForecastSummary struct {
	BriefPhrase   string `json:"briefPhrase,omitempty"`
	BriefPhrase60 string `json:"briefPhrase60,omitempty"`
	IconCode      int    `json:"iconCode"`
	LongPhrase    string `json:"longPhrase,omitempty"`
	ShortPhrase   string `json:"shortPhrase,omitempty"`
}
