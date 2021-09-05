package entities

type DailyIndex struct {
	Ascending     bool    `json:"ascending"`
	Category      string  `json:"category,omitempty"`
	CategoryValue int     `json:"categoryValue"`
	DateTime      string  `json:"dateTime,omitempty"`
	Description   string  `json:"description,omitempty"`
	IndexId       int     `json:"indexId"`
	IndexName     string  `json:"indexName,omitempty"`
	Value         float64 `json:"value"`
}
