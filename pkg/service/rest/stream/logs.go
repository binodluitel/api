package stream

import (
	"context"
	"fmt"
	"io"

	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Stream) StreamLogs(ctx context.Context, request *models.StreamRequest) (io.ReadCloser, error) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	pods, err := s.k8sClient.CoreV1().Pods(request.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed listing pods: %w", err)
	}
	if len(pods.Items) == 0 {
		return nil, fmt.Errorf("no pods found in namespace %s", request.Namespace)
	}
	// For now, we are just returning logs for the first pod found
	// TODO: Add support for selecting a specific pod or all pods
	podName := pods.Items[0].Name
	logs := s.k8sClient.CoreV1().Pods(request.Namespace).GetLogs(podName, &corev1.PodLogOptions{
		Container:    request.Container,
		Follow:       request.Follow,
		SinceSeconds: request.SinceSeconds,
		Timestamps:   true,
	})
	return logs.Stream(ctx)
}
