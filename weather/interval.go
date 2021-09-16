package azweather

type Interval struct {
	CloudCover        int     `json:"cloudCover"`
	Color             *Color  `json:"color,omitempty"`
	Dbz               float64 `json:"dbz"`
	IconCode          int     `json:"iconCode"`
	Minute            int     `json:"minute"`
	PrecipitationType string  `json:"precipitationType,omitempty"`
	ShortPhrase       string  `json:"shortPhrase,omitempty"`
	SimplifiedColor   *Color  `json:"simplifiedColor,omitempty"`
	StartTime         string  `json:"startTime,omitempty"`
	Threshold         string  `json:"threshold,omitempty"`
}
