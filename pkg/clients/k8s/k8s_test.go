package k8s_test

import (
	"testing"

	"github.com/binodluitel/api/pkg/clients/k8s"
	"github.com/binodluitel/api/pkg/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "k8s client test suite")
}

var _ = Describe("kubeConfig", func() {
	It("should return an error if failed to get kube config", func() {
		cfg := &config.Config{
			KubeConfigPath: "~/invalid/kubeconfig", // hope user does not have this file in their home directory
		}
		_, err := k8s.New(cfg)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("failed to get kubernetes config"))
	})
	It("should create a client successfully with valid kube config", func() {
		cfg := &config.Config{
			KubeConfigPath: "../../fixtures/test_k8s_config",
		}
		client, err := k8s.New(cfg)
		Expect(err).NotTo(HaveOccurred())
		Expect(client).NotTo(BeNil())
	})
})
