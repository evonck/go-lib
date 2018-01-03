package models

import (
	"net/http"
)

// SyncHTTPResponse model for async response
type SyncHTTPResponse struct {
	URL      string
	Response *http.Response
}

// AppErr model
type AppErr struct {
	Code          int      `json:"code"`
	StatusCode    string   `json:"statusCode"`
	Message       string   `json:"message"`
	Detail        string   `json:"detail"`
	Params        []string `json:"params"`
	CorrelationID string   `json:"correlationId"`
	Ts            int64    `json:"ts"`
}
