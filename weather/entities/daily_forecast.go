package entities

type DailyForecast struct {
	AirAndPollen             []AirAndPollen   `json:"airAndPollen,omitempty"`
	Date                     string           `json:"date,omitempty"`
	Day                      DayOrNight       `json:"day,omitempty"`
	DegreeDaySummary         DegreeDaySummary `json:"degreeDaySummary,omitempty"`
	HoursOfSun               float64          `json:"hoursOfSun,omitempty"`
	Night                    DayOrNight       `json:"night,omitempty"`
	RealFeelTemperature      WeatherUnitRange `json:"realFeelTemperature,omitempty"`
	RealFeelTemperatureShade WeatherUnitRange `json:"realFeelTemperatureShade,omitempty"`
	Sources                  []string         `json:"sources,omitempty"`
	Temperature              WeatherUnitRange `json:"temperature,omitempty"`
}
