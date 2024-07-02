package config_test

import (
	"testing"

	"github.com/binodluitel/api/pkg/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Config", func() {
	Describe("MustGet", func() {
		Context("when config is initialized", func() {
			It("should return the config", func() {
				cfg := config.MustGet()
				Expect(cfg).ToNot(BeNil())
			})
		})
	})

	Describe("Get", func() {
		Context("when config is initialized", func() {
			It("should return the config and no error", func() {
				cfg, err := config.Get()
				Expect(err).To(BeNil())
				Expect(cfg).ToNot(BeNil())
			})
		})
	})
})
