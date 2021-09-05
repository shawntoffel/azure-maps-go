package entities

type CurrentConditions struct {
	ApparentTemperature            *WeatherUnit          `json:"apparentTemperature"`
	Ceiling                        *WeatherUnit          `json:"ceiling"`
	CloudCover                     int                   `json:"cloudCover"`
	DateTime                       string                `json:"dateTime"`
	DewPoint                       *WeatherUnit          `json:"dewPoint"`
	HasPrecipitation               bool                  `json:"hasPrecipitation"`
	IconCode                       int                   `json:"iconCode"`
	IsDayTime                      bool                  `json:"isDayTime"`
	ObstructionsToVisibility       string                `json:"obstructionsToVisibility"`
	Past24HourTemperatureDeparture *WeatherUnit          `json:"past24HourTemperatureDeparture"`
	Phrase                         string                `json:"phrase"`
	PrecipitationSummary           *PrecipitationSummary `json:"precipitationSummary"`
	Pressure                       *WeatherUnit          `json:"pressure"`
	PressureTendency               *PressureTendency     `json:"pressureTendency"`
	RealFeelTemperature            *WeatherUnit          `json:"realFeelTemperature"`
	RealFeelTemperatureShade       *WeatherUnit          `json:"realFeelTemperatureShade"`
	RelativeHumidity               int                   `json:"relativeHumidity"`
	Temperature                    *WeatherUnit          `json:"temperature"`
	TemperatureSummary             *TemperatureSummary   `json:"temperatureSummary"`
	UVIndex                        int                   `json:"uvIndex"`
	UVIndexPhrase                  string                `json:"uvIndexPhrase"`
	Visibility                     *WeatherUnit          `json:"visibility"`
	WetBulbTemperature             *WeatherUnit          `json:"wetBulbTemperature"`
	Wind                           *Wind                 `json:"wind"`
	WindChillTemperature           *WeatherUnit          `json:"windChillTemperature"`
	WindGust                       *Wind                 `json:"windGust"`
}
