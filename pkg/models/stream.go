package models

// StreamRequest is request model for streaming
type StreamRequest struct {
	Follow       bool   `json:"follow"`
	SinceSeconds *int64 `json:"since_seconds"`
	Container    string `json:"container"`
	Namespace    string `json:"namespace"`
}
