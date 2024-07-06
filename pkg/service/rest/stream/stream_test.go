package stream_test

import (
	"testing"

	"github.com/binodluitel/api/pkg/config"
	"github.com/binodluitel/api/pkg/service/rest/stream"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "REST stream service test suite")
}

var _ = Describe("Stream", func() {
	Describe("New", func() {
		It("should return a new Stream instance", func() {
			service, err := stream.New(&config.Config{
				KubeConfigPath: "../../../fixtures/test_k8s_config",
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(service).ToNot(BeNil())
			Expect(service).To(BeAssignableToTypeOf(&stream.Stream{}))
		})
		It("should error when the configuration is not valid", func() {
			service, err := stream.New(&config.Config{
				KubeConfigPath: "~/invalid/kubeconfig",
			})
			Expect(err).To(HaveOccurred())
			Expect(service).To(BeNil())
		})
	})
})
