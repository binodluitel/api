package models

// StreamRequest is request model for streaming
type StreamRequest struct {
	Follow       bool   `json:"follow"`
	SinceSeconds *int64 `json:"since_seconds,omitempty"`
	Container    string `json:"container,omitempty"`
	Namespace    string `json:"namespace,omitempty"`
}
