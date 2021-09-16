package weather

import (
	"fmt"
)

type ODataErrorResponse struct {
	Error ODataError `json:"error,omitempty"`
}

type ODataError struct {
	Code    string       `json:"code,omitempty"`
	Details []ODataError `json:"details,omitempty"`
	Message string       `json:"message,omitempty"`
	Target  string       `json:"target,omitempty"`
}

func (o ODataError) Error() string {
	errorMessage := fmt.Sprintf("code: %s, message: %s", o.Code, o.Message)
	if o.Target != "" {
		errorMessage += ", target: " + o.Target
	}
	if len(o.Details) > 0 {
		errorMessage += fmt.Sprintf(", details: %#v", o.Details)
	}

	return errorMessage
}
