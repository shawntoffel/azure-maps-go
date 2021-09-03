package entities

type HourlyForecast struct {
	Ceiling                  WeatherUnit `json:"ceiling,omitempty"`
	CloudCover               int         `json:"cloudCover,omitempty"`
	Date                     string      `json:"date,omitempty"`
	DewPoint                 WeatherUnit `json:"dewPoint,omitempty"`
	HasPrecipitation         bool        `json:"hasPrecipitation,omitempty"`
	Ice                      WeatherUnit `json:"ice,omitempty"`
	IceProbability           int         `json:"iceProbability,omitempty"`
	IconCode                 int         `json:"iconCode,omitempty"`
	IconPhrase               string      `json:"iconPhrase,omitempty"`
	IsDaylight               bool        `json:"isDaylight,omitempty"`
	PrecipitationProbability int         `json:"precipitationProbability,omitempty"`
	Rain                     WeatherUnit `json:"rain,omitempty"`
	RainProbability          int         `json:"rainProbability,omitempty"`
	RealFeelTemperature      WeatherUnit `json:"realFeelTemperature,omitempty"`
	RelativeHumidity         int         `json:"relativeHumidity,omitempty"`
	Snow                     WeatherUnit `json:"snow,omitempty"`
	SnowProbability          int         `json:"snowProbability,omitempty"`
	Temperature              WeatherUnit `json:"temperature,omitempty"`
	TotalLiquid              WeatherUnit `json:"totalLiquid,omitempty"`
	UvIndex                  int         `json:"uvIndex,omitempty"`
	UvIndexPhrase            string      `json:"uvIndexPhrase,omitempty"`
	Visibility               WeatherUnit `json:"visibility,omitempty"`
	WetBulbTemperature       WeatherUnit `json:"wetBulbTemperature,omitempty"`
	Wind                     Wind        `json:"wind,omitempty"`
	WindGust                 Wind        `json:"windGust,omitempty"`
}
