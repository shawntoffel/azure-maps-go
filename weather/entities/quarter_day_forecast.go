package entities

type QuarterDayForecast struct {
	CloudCover               int              `json:"cloudCover,omitempty"`
	Date                     string           `json:"date,omitempty"`
	DewPoint                 WeatherUnit      `json:"dewPoint,omitempty"`
	EffectiveDate            string           `json:"effectiveDate,omitempty"`
	HasPrecipitation         bool             `json:"hasPrecipitation,omitempty"`
	Ice                      WeatherUnit      `json:"ice,omitempty"`
	IconCode                 int              `json:"iconCode,omitempty"`
	IconPhrase               string           `json:"iconPhrase,omitempty"`
	Phrase                   string           `json:"phrase,omitempty"`
	PrecipitationIntensity   string           `json:"precipitationIntensity,omitempty"`
	PrecipitationProbability int              `json:"precipitationProbability,omitempty"`
	PrecipitationType        string           `json:"precipitationType,omitempty"`
	Quarter                  Quarter          `json:"quarter,omitempty"`
	Rain                     WeatherUnit      `json:"rain,omitempty"`
	RealFeelTemperature      WeatherUnitRange `json:"realFeelTemperature,omitempty"`
	RelativeHumidity         int              `json:"relativeHumidity,omitempty"`
	Snow                     WeatherUnit      `json:"snow,omitempty"`
	Temperature              WeatherUnitRange `json:"temperature,omitempty"`
	ThunderstormProbability  int              `json:"thunderstormProbability,omitempty"`
	TotalLiquid              WeatherUnit      `json:"totalLiquid,omitempty"`
	Visibility               WeatherUnit      `json:"visibility,omitempty"`
	Wind                     Wind             `json:"wind,omitempty"`
	WindGust                 Wind             `json:"windGust,omitempty"`
}
