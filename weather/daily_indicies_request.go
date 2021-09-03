package weather

import (
	"net/url"
	"strconv"
)

type DailyIndicesRequestOptions struct {
	Duration     *int   `json:"duration,omitempty"`
	IndexGroupId *int   `json:"indexGroupId,omitempty"`
	IndexId      *int   `json:"indexId,omitempty"`
	Language     string `json:"language,omitempty"`
}

func (ccr DailyIndicesRequestOptions) Encode() string {
	q := url.Values{}

	if ccr.Duration != nil {
		q.Add("duration", strconv.Itoa(*ccr.Duration))
	}

	if ccr.IndexGroupId != nil {
		q.Add("indexGroupId", strconv.Itoa(*ccr.IndexGroupId))
	}

	if ccr.IndexId != nil {
		q.Add("indexId", strconv.Itoa(*ccr.IndexId))
	}

	if ccr.Language != "" {
		q.Add("language", ccr.Language)
	}

	return q.Encode()
}
