package azweather

type LocalSource struct {
	ID          int    `json:"id"`
	Name        string `json:"name,omitempty"`
	WeatherCode string `json:"weatherCode,omitempty"`
}
