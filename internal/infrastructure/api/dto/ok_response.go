package dto

import (
	"time"
)

type OKResponse struct {
	Status    int         `json:"status"`
	Body      interface{} `json:"body"`
	IssuedAt  time.Time   `json:"issuedAt"`
	RequestID any         `json:"requestID"`
}
