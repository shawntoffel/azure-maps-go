package entities

type LocalSource struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	WeatherCode string `json:"weatherCode,omitempty"`
}
