package integration_test

import (
	"net/http"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Custom HTTP Headers", func() {

	var tmp string

	BeforeEach(func() {
		var err error
		tmp, err = os.MkdirTemp("", "custom-http-headers-test")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(tmp)
	})

	Context("when the flag is not set", func() {
		It("starts successfully and does not set any custom headers", func() {
			resp, err := http.Get(atcURL)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.Header.Get("X-Custom-Header")).To(Equal(""))
		})
	})

	Context("when the file does not exist", func() {
		It("returns an error on startup", func() {
			err := cmd.Server.CustomHTTPHeaders.UnmarshalFlag("/nonexistent/headers.yml")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Failed to open"))
		})
	})

	Context("when the file contains invalid YAML", func() {
		It("returns an error on startup", func() {
			path := filepath.Join(tmp, "headers.yml")
			Expect(os.WriteFile(path, []byte("X-Custom-Header; custom-value\n"), 0644)).To(Succeed())

			err := cmd.Server.CustomHTTPHeaders.UnmarshalFlag(path)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Failed to parse"))
		})
	})

	Context("when the file contains an invalid header name", func() {
		It("returns an error on startup", func() {
			path := filepath.Join(tmp, "headers.yml")
			Expect(os.WriteFile(path, []byte("X Custom Header: custom-value\n"), 0644)).To(Succeed())

			err := cmd.Server.CustomHTTPHeaders.UnmarshalFlag(path)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Invalid header name"))
		})
	})

	Context("when the file is valid", func() {
		BeforeEach(func() {
			path := filepath.Join(tmp, "headers.yml")
			Expect(os.WriteFile(path, []byte("X-Custom-Header: \"some-custom-value\"\n"), 0644)).To(Succeed())

			err := cmd.Server.CustomHTTPHeaders.UnmarshalFlag(path)
			Expect(err).ToNot(HaveOccurred())
		})

		It("sets the custom header on responses", func() {
			resp, err := http.Get(atcURL)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.Header.Get("X-Custom-Header")).To(Equal("some-custom-value"))
		})

		It("overrides a hardcoded header", func() {
			path := filepath.Join(tmp, "override.yml")
			Expect(os.WriteFile(path, []byte("X-Content-Type-Options: \"custom-value\"\n"), 0644)).To(Succeed())

			err := cmd.Server.CustomHTTPHeaders.UnmarshalFlag(path)
			Expect(err).ToNot(HaveOccurred())

			resp, err := http.Get(atcURL)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.Header.Get("X-Content-Type-Options")).To(Equal("custom-value"))
		})

		It("preserves default security headers", func() {
			resp, err := http.Get(atcURL)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.Header.Get("X-Frame-Options")).To(Equal("deny"))
			Expect(resp.Header.Get("X-Download-Options")).To(Equal("noopen"))
			Expect(resp.Header.Get("Cache-Control")).To(Equal("no-store, private"))
		})

		Context("when the file is JSON", func() {
			BeforeEach(func() {
				path := filepath.Join(tmp, "headers.json")
				Expect(os.WriteFile(path, []byte(`{"X-Custom-Header": "some-custom-value"}`), 0644)).To(Succeed())

				err := cmd.Server.CustomHTTPHeaders.UnmarshalFlag(path)
				Expect(err).ToNot(HaveOccurred())
			})

			It("sets the custom header on responses", func() {
				resp, err := http.Get(atcURL)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Header.Get("X-Custom-Header")).To(Equal("some-custom-value"))
			})
		})
	})
})
