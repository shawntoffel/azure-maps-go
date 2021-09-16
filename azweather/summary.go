package azweather

type DailyForecastSummary struct {
	Category  string `json:"category,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	Phrase    string `json:"phrase,omitempty"`
	Severity  int    `json:"severity"`
	StartDate string `json:"startDate,omitempty"`
}

type DegreeDaySummary struct {
	Cooling *WeatherUnit `json:"cooling,omitempty"`
	Heating *WeatherUnit `json:"heating,omitempty"`
}

type IntervalSummary struct {
	BriefPhrase  string `json:"briefPhrase,omitempty"`
	EndMinute    int    `json:"endMinute"`
	IconCode     int    `json:"iconCode"`
	LongPhrase   string `json:"longPhrase,omitempty"`
	ShortPhrase  string `json:"shortPhrase,omitempty"`
	StartMinute  int    `json:"startMinute"`
	TotalMinutes int    `json:"totalMinutes"`
}

type MinuteForecastSummary struct {
	BriefPhrase   string `json:"briefPhrase,omitempty"`
	BriefPhrase60 string `json:"briefPhrase60,omitempty"`
	IconCode      int    `json:"iconCode"`
	LongPhrase    string `json:"longPhrase,omitempty"`
	ShortPhrase   string `json:"shortPhrase,omitempty"`
}

type PrecipitationSummary struct {
	Past12Hours *WeatherUnit `json:"past12Hours,omitempty"`
	Past18Hours *WeatherUnit `json:"past18Hours,omitempty"`
	Past24Hours *WeatherUnit `json:"past24Hours,omitempty"`
	Past3Hours  *WeatherUnit `json:"past3Hours,omitempty"`
	Past6Hours  *WeatherUnit `json:"past6Hours,omitempty"`
	Past9Hours  *WeatherUnit `json:"past9Hours,omitempty"`
	PastHour    *WeatherUnit `json:"pastHour,omitempty"`
}

type TemperatureSummary struct {
	Past12Hours *WeatherUnitRange `json:"past12Hours,omitempty"`
	Past24Hours *WeatherUnitRange `json:"past24Hours,omitempty"`
	Past6Hours  *WeatherUnitRange `json:"past6Hours,omitempty"`
}

type WeatherAlongRouteSummary struct {
	Hazards  *Hazards `json:"hazards,omitempty"`
	IconCode int      `json:"iconCode"`
}
