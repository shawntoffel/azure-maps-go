package azweather

type LatestStatus struct {
	English   LatestStatusKeyword `json:"english,omitempty"`
	Localized string              `json:"localized,omitempty"`
}
