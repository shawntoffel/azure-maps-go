package entities

type AirAndPollen struct {
	Category      string `json:"category,omitempty"`
	CategoryValue int    `json:"categoryValue,omitempty"`
	Name          string `json:"name,omitempty"`
	Type          string `json:"type,omitempty"`
	Value         int    `json:"value,omitempty"`
}
