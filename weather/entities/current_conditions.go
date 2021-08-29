package entities

type CurrentConditions struct {
	ApparentTemperature            WeatherUnit          `json:"apparentTemperature,omitempty"`
	Ceiling                        WeatherUnit          `json:"ceiling,omitempty"`
	CloudClover                    int                  `json:"cloudClover,omitempty"`
	DateTime                       string               `json:"dateTime,omitempty"`
	DewPoint                       WeatherUnit          `json:"dewPoint,omitempty"`
	HasPrecipitation               bool                 `json:"hasPrecipitation,omitempty"`
	IconCode                       int                  `json:"iconCode,omitempty"`
	IsDayTime                      bool                 `json:"isDayTime,omitempty"`
	ObstructionsToVisibility       string               `json:"obstructionsToVisibility,omitempty"`
	Past24HourTemperatureDeparture WeatherUnit          `json:"past24HourTemperatureDeparture,omitempty"`
	Phrase                         string               `json:"phrase,omitempty"`
	PrecipitationSummary           PrecipitationSummary `json:"precipitationSummary,omitempty"`
	Pressure                       WeatherUnit          `json:"pressure,omitempty"`
	PressureTendency               PressureTendency     `json:"pressureTendency,omitempty"`
	RealFeelTemperature            WeatherUnit          `json:"realFeelTemperature,omitempty"`
	RealFeelTemperatureShade       WeatherUnit          `json:"realFeelTemperatureShade,omitempty"`
	RelativeHumidity               int                  `json:"relativeHumidity,omitempty"`
	Temperature                    WeatherUnit          `json:"temperature,omitempty"`
	TemeratureSummary              TemeratureSummary    `json:"temeratureSummary,omitempty"`
	UVIndex                        int                  `json:"uvIndex,omitempty"`
	UVIndexPhrase                  string               `json:"uvIndexPhrase,omitempty"`
	Visibility                     WeatherUnit          `json:"visibility,omitempty"`
	WetBulbTemperature             WeatherUnit          `json:"wetBulbTemperature,omitempty"`
	Wind                           Wind                 `json:"wind,omitempty"`
	WindChillTemperature           WeatherUnit          `json:"windChillTemperature,omitempty"`
	WindGust                       Wind                 `json:"windGust,omitempty"`
}
