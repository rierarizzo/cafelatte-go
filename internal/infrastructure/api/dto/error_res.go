package dto

import (
	"time"
)

type ErrorResponse struct {
	Status    int       `json:"status"`
	ErrorType string    `json:"errorType"`
	ErrorMsg  string    `json:"errorMsg"`
	IssuedAt  time.Time `json:"issuedAt"`
	RequestID any       `json:"requestID"`
}
