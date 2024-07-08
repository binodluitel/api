package pods

import (
	"context"
	"fmt"
	"io"

	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models/pods"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
)

func (p *Pods) GetLogs(ctx context.Context, request *pods.Logs) (io.ReadCloser, error) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	logger.Info("Streaming pod logs", zap.String("pod_name", request.PodName))
	r := p.k8sClient.CoreV1().Pods(request.Namespace).GetLogs(request.PodName, &corev1.PodLogOptions{
		Container:  request.Container,
		Follow:     request.Follow,
		Timestamps: true,
	})
	logs, err := r.Stream(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs for pod %q: %w", request.PodName, err)
	}
	return logs, nil
}
