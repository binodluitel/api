package pods_test

import (
	"testing"

	"github.com/binodluitel/api/pkg/config"
	"github.com/binodluitel/api/pkg/service/rest/pods"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "REST pods service test suite")
}

var _ = Describe("Pods", func() {
	Describe("New", func() {
		It("should return a new Pods service instance", func() {
			service, err := pods.New(&config.Config{
				KubeConfigPath: "../../../fixtures/test_k8s_config",
			})
			Expect(err).ToNot(HaveOccurred())
			Expect(service).ToNot(BeNil())
			Expect(service).To(BeAssignableToTypeOf(&pods.Pods{}))
		})
		It("should error when the configuration is not valid", func() {
			service, err := pods.New(&config.Config{
				KubeConfigPath: "~/invalid/kubeconfig",
			})
			Expect(err).To(HaveOccurred())
			Expect(service).To(BeNil())
		})
	})
})
