package log_test

import (
	"context"
	"testing"

	"github.com/binodluitel/api/pkg/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSuiteLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Log Suite")
}

var _ = Describe("Log", func() {
	Describe("New", func() {
		It("should initialize a new logger", func() {
			logger, err := log.New(context.Background())
			Expect(err).To(BeNil())
			Expect(logger).NotTo(BeNil())
		})
	})

	Describe("Get", func() {
		It("should return an initialized logger", func() {
			ctx := context.Background()
			_, logger := log.Get(ctx)
			Expect(logger).NotTo(BeNil())
		})
	})

	Describe("WithTrace", func() {
		It("should return a logger with trace span initialized", func() {
			ctx := context.Background()
			_, logger := log.WithTrace(ctx, "operation")
			Expect(logger).NotTo(BeNil())
		})
	})

	Describe("Sync", func() {
		It("should synchronize the logger", func() {
			logger, err := log.New(context.Background())
			Expect(err).To(BeNil())
			Expect(logger).NotTo(BeNil())

			Expect(func() { logger.Sync() }).NotTo(Panic())
		})
	})

	Describe("With", func() {
		It("should return a new logger with additional fields", func() {
			logger, err := log.New(context.Background())
			Expect(err).To(BeNil())
			Expect(logger).NotTo(BeNil())

			newLogger := logger.With()
			Expect(newLogger).NotTo(BeNil())
		})
	})
})
