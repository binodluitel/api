package k8s

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/binodluitel/api/pkg/config"
	"github.com/mitchellh/go-homedir"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func New(cfg *config.Config) (*kubernetes.Clientset, error) {
	config, err := kubeConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to get kubernetes config: %w", err)
	}
	return kubernetes.NewForConfig(config)
}

// kubeConfig returns a Kubernetes client configuration based on kubeconfig path or
// in-cluster configuration if kubeconfig path is empty.
func kubeConfig(cfg *config.Config) (*rest.Config, error) {
	kubeConfigPath := cfg.KubeConfigPath
	if kubeConfigPath != "" && strings.HasPrefix(kubeConfigPath, "~/") {
		home, err := homedir.Dir()
		if err != nil {
			return nil, fmt.Errorf("failed to get kube config from path %q: %w", kubeConfigPath, err)
		}
		kubeConfigPath = filepath.Join(home, kubeConfigPath[2:])
	}
	return clientcmd.BuildConfigFromFlags("", kubeConfigPath)
}
