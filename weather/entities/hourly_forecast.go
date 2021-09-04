package entities

type HourlyForecast struct {
	Ceiling                  *WeatherUnit `json:"ceiling,omitempty"`
	CloudCover               int          `json:"cloudCover"`
	Date                     string       `json:"date,omitempty"`
	DewPoint                 *WeatherUnit `json:"dewPoint,omitempty"`
	HasPrecipitation         bool         `json:"hasPrecipitation"`
	Ice                      *WeatherUnit `json:"ice,omitempty"`
	IceProbability           int          `json:"iceProbability"`
	IconCode                 int          `json:"iconCode"`
	IconPhrase               string       `json:"iconPhrase,omitempty"`
	IsDaylight               bool         `json:"isDaylight"`
	PrecipitationProbability int          `json:"precipitationProbability"`
	Rain                     *WeatherUnit `json:"rain,omitempty"`
	RainProbability          int          `json:"rainProbability"`
	RealFeelTemperature      *WeatherUnit `json:"realFeelTemperature,omitempty"`
	RelativeHumidity         int          `json:"relativeHumidity"`
	Snow                     *WeatherUnit `json:"snow,omitempty"`
	SnowProbability          int          `json:"snowProbability"`
	Temperature              *WeatherUnit `json:"temperature,omitempty"`
	TotalLiquid              *WeatherUnit `json:"totalLiquid,omitempty"`
	UvIndex                  int          `json:"uvIndex"`
	UvIndexPhrase            string       `json:"uvIndexPhrase,omitempty"`
	Visibility               *WeatherUnit `json:"visibility,omitempty"`
	WetBulbTemperature       *WeatherUnit `json:"wetBulbTemperature,omitempty"`
	Wind                     *Wind        `json:"wind,omitempty"`
	WindGust                 *Wind        `json:"windGust,omitempty"`
}
