package pods

// Logs is for getting pod logs
type Logs struct {
	Follow    bool   `json:"follow"`
	Container string `json:"container,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	PodName   string `json:"pod_name" binding:"required"`
}
