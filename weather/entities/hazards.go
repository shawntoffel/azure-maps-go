package entities

type Hazards struct {
	HazardDetails  []HazardDetail `json:"hazardDetails,omitempty"`
	MaxHazardIndex int            `json:"maxHazardIndex"`
}
