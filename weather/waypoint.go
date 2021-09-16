package azweather

type Waypoint struct {
	CloudCover     int                             `json:"cloudCover"`
	Hazards        *Hazards                        `json:"hazards,omitempty"`
	IconCode       int                             `json:"iconCode"`
	IsDayTime      bool                            `json:"isDayTime"`
	LightningCount int                             `json:"lightningCount"`
	Notifications  []Notification                  `json:"notifications"`
	Precipitation  *WeatherAlongRoutePrecipitation `json:"precipitation,omitempty"`
	ShortPhrase    string                          `json:"shortPhrase,omitempty"`
	SunGlare       *SunGlare                       `json:"sunGlare,omitempty"`
	Temperature    *WeatherUnit                    `json:"temperature,omitempty"`
	Wind           *Wind                           `json:"wind,omitempty"`
	WindGust       *Wind                           `json:"windGust,omitempty"`
}
