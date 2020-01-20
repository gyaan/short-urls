package models

import "encoding/json"

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	Retry        bool   `json:"retry"`
}

func (e *ErrorResponse) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}
