package azweather

type DailyForecast struct {
	AirAndPollen             []AirAndPollen    `json:"airAndPollen,omitempty"`
	Date                     string            `json:"date,omitempty"`
	Day                      *DayOrNight       `json:"day,omitempty"`
	DegreeDaySummary         *DegreeDaySummary `json:"degreeDaySummary,omitempty"`
	HoursOfSun               float64           `json:"hoursOfSun"`
	Night                    *DayOrNight       `json:"night,omitempty"`
	RealFeelTemperature      *WeatherUnitRange `json:"realFeelTemperature,omitempty"`
	RealFeelTemperatureShade *WeatherUnitRange `json:"realFeelTemperatureShade,omitempty"`
	Sources                  []string          `json:"sources,omitempty"`
	Temperature              *WeatherUnitRange `json:"temperature,omitempty"`
}

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

type QuarterDayForecast struct {
	CloudCover               int               `json:"cloudCover"`
	Date                     string            `json:"date,omitempty"`
	DewPoint                 *WeatherUnit      `json:"dewPoint,omitempty"`
	EffectiveDate            string            `json:"effectiveDate,omitempty"`
	HasPrecipitation         bool              `json:"hasPrecipitation"`
	Ice                      *WeatherUnit      `json:"ice,omitempty"`
	IconCode                 int               `json:"iconCode"`
	IconPhrase               string            `json:"iconPhrase,omitempty"`
	Phrase                   string            `json:"phrase,omitempty"`
	PrecipitationIntensity   string            `json:"precipitationIntensity,omitempty"`
	PrecipitationProbability int               `json:"precipitationProbability"`
	PrecipitationType        string            `json:"precipitationType,omitempty"`
	Quarter                  Quarter           `json:"quarter"`
	Rain                     *WeatherUnit      `json:"rain,omitempty"`
	RealFeelTemperature      *WeatherUnitRange `json:"realFeelTemperature,omitempty"`
	RelativeHumidity         int               `json:"relativeHumidity"`
	Snow                     *WeatherUnit      `json:"snow,omitempty"`
	Temperature              *WeatherUnitRange `json:"temperature,omitempty"`
	ThunderstormProbability  int               `json:"thunderstormProbability"`
	TotalLiquid              *WeatherUnit      `json:"totalLiquid,omitempty"`
	Visibility               *WeatherUnit      `json:"visibility,omitempty"`
	Wind                     *Wind             `json:"wind,omitempty"`
	WindGust                 *Wind             `json:"windGust,omitempty"`
}
