package pods

// Logs is for getting pod logs
type Logs struct {
	Follow       bool   `json:"follow"`
	SinceSeconds *int64 `json:"since_seconds,omitempty"`
	Container    string `json:"container,omitempty"`
	Namespace    string `json:"namespace,omitempty"`
	PodName      string `json:"pod_name" binding:"required"`
}
