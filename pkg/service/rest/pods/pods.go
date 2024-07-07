package pods

import (
	"fmt"

	"github.com/binodluitel/api/pkg/clients/k8s"
	"github.com/binodluitel/api/pkg/config"
	svcdef "github.com/binodluitel/api/pkg/service/definitions"
	"k8s.io/client-go/kubernetes"
)

// Pods defines service instance for k8s pods
type Pods struct {
	k8sClient *kubernetes.Clientset
}

// New creates and returns a new user service instance
func New(cfg *config.Config) (svcdef.PodsService, error) {
	k8sClient, err := k8s.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed initializing k8s client: %w", err)
	}
	return &Pods{k8sClient: k8sClient}, nil
}
