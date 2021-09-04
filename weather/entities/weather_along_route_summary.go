package entities

type WeatherAlongRouteSummary struct {
	Hazards  *Hazards `json:"hazards,omitempty"`
	IconCode int      `json:"iconCode"`
}
