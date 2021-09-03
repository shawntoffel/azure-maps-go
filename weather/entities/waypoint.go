package entities

type Waypoint struct {
	CloudCover     int                            `json:"cloudCover,omitempty"`
	Hazards        Hazards                        `json:"hazards,omitempty"`
	IconCode       int                            `json:"iconCode,omitempty"`
	IsDayTime      bool                           `json:"isDayTime,omitempty"`
	LightningCount int                            `json:"lightningCount,omitempty"`
	Notifications  []Notification                 `json:"notifications,omitempty"`
	Precipitation  WeatherAlongRoutePrecipitation `json:"precipitation,omitempty"`
	ShortPhrase    string                         `json:"shortPhrase,omitempty"`
	SunGlare       SunGlare                       `json:"sunGlare,omitempty"`
	Temperature    WeatherUnit                    `json:"temperature,omitempty"`
	Wind           Wind                           `json:"wind,omitempty"`
	WindGust       Wind                           `json:"windGust,omitempty"`
}
