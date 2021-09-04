package entities

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
