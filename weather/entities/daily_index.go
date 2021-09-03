package entities

type DailyIndex struct {
	Ascending     bool    `json:"ascending,omitempty"`
	Category      string  `json:"category,omitempty"`
	CategoryValue int     `json:"categoryValue,omitempty"`
	DateTime      string  `json:"dateTime,omitempty"`
	Description   string  `json:"description,omitempty"`
	IndexId       int     `json:"indexId,omitempty"`
	IndexName     string  `json:"indexName,omitempty"`
	Value         float64 `json:"value,omitempty"`
}
