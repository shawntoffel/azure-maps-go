package weather

type DayOrNight struct {
	CloudCover               int          `json:"cloudCover"`
	HasPrecipitation         bool         `json:"hasPrecipitation"`
	HoursOfIce               float64      `json:"hoursOfIce"`
	HoursOfPrecipitation     float64      `json:"hoursOfPrecipitation"`
	HoursOfRain              float64      `json:"hoursOfRain"`
	HoursOfSnow              float64      `json:"hoursOfSnow"`
	Ice                      *WeatherUnit `json:"ice,omitempty"`
	IceProbability           int          `json:"iceProbability"`
	IconCode                 int          `json:"iconCode"`
	IconPhrase               string       `json:"iconPhrase,omitempty"`
	LocalSource              *LocalSource `json:"localSource,omitempty"`
	LongPhrase               string       `json:"longPhrase,omitempty"`
	PrecipitationIntensity   string       `json:"precipitationIntensity,omitempty"`
	PrecipitationProbability int          `json:"precipitationProbability"`
	PrecipitationType        string       `json:"precipitationType,omitempty"`
	Rain                     *WeatherUnit `json:"rain,omitempty"`
	RainProbability          int          `json:"rainProbability"`
	ShortPhrase              string       `json:"shortPhrase,omitempty"`
	Snow                     *WeatherUnit `json:"snow,omitempty"`
	SnowProbability          int          `json:"snowProbability"`
	ThunderstormProbability  int          `json:"thunderstormProbability"`
	TotalLiquid              *WeatherUnit `json:"totalLiquid,omitempty"`
	Wind                     *Wind        `json:"wind,omitempty"`
	WindGust                 *Wind        `json:"windGust,omitempty"`
}
