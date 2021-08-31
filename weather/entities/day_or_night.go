package entities

type DayOrNight struct {
	CloudCover               int         `json:"cloudCover,omitempty"`
	HasPrecipitation         bool        `json:"hasPrecipitation,omitempty"`
	HoursOfIce               float64     `json:"hoursOfIce,omitempty"`
	HoursOfPrecipitation     float64     `json:"hoursOfPrecipitation,omitempty"`
	HoursOfRain              float64     `json:"hoursOfRain,omitempty"`
	HoursOfSnow              float64     `json:"hoursOfSnow,omitempty"`
	Ice                      WeatherUnit `json:"ice,omitempty"`
	IceProbability           int         `json:"iceProbability,omitempty"`
	IconCode                 int         `json:"iconCode,omitempty"`
	IconPhrase               string      `json:"iconPhrase,omitempty"`
	LocalSource              LocalSource `json:"localSource,omitempty"`
	LongPhrase               string      `json:"longPhrase,omitempty"`
	PrecipitationIntensity   string      `json:"precipitationIntensity,omitempty"`
	PrecipitationProbability int         `json:"precipitationProbability,omitempty"`
	PrecipitationType        string      `json:"precipitationType,omitempty"`
	Rain                     WeatherUnit `json:"rain,omitempty"`
	RainProbability          int         `json:"rainProbability,omitempty"`
	ShortPhrase              string      `json:"shortPhrase,omitempty"`
	Snow                     WeatherUnit `json:"snow,omitempty"`
	SnowProbability          int         `json:"snowProbability,omitempty"`
	ThunderstormProbability  int         `json:"thunderstormProbability,omitempty"`
	TotalLiquid              WeatherUnit `json:"totalLiquid,omitempty"`
	Wind                     Wind        `json:"wind,omitempty"`
	WindGust                 Wind        `json:"windGust,omitempty"`
}
