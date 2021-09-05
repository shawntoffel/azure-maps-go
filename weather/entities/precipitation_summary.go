package entities

type PrecipitationSummary struct {
	Past12Hours *WeatherUnit `json:"past12Hours,omitempty"`
	Past18Hours *WeatherUnit `json:"past18Hours,omitempty"`
	Past24Hours *WeatherUnit `json:"past24Hours,omitempty"`
	Past3Hours  *WeatherUnit `json:"past3Hours,omitempty"`
	Past6Hours  *WeatherUnit `json:"past6Hours,omitempty"`
	Past9Hours  *WeatherUnit `json:"past9Hours,omitempty"`
	PastHour    *WeatherUnit `json:"pastHour,omitempty"`
}
