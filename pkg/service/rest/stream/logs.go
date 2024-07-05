package stream

import (
	"context"

	"github.com/binodluitel/api/pkg/log"
	"github.com/binodluitel/api/pkg/models"
	"k8s.io/apimachinery/pkg/util/rand"
)

func (s *Stream) StreamLogs(ctx context.Context, request *models.StreamRequest) (*string, error) {
	_, logger := log.Get(ctx)
	defer logger.Sync()
	// TODO: Implement the StreamLogs function that actually streams logs.
	giberish := generateGibberish(10)
	return &giberish, nil
}

// generateGibberish generates a string of random characters of a given length.
func generateGibberish(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b) + "\n"
}
