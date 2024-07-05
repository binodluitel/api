package models

import (
	"time"
)

// StreamRequest is request model for streaming
type StreamRequest struct {
	ID     string    `json:"id"`
	Follow bool      `json:"follow"`
	Tail   int       `json:"tail"`
	Limit  int64     `json:"limit"`
	Since  time.Time `json:"since"`
}
