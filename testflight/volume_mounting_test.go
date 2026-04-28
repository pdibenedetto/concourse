package testflight_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("A job with nested volume mounts", func() {
	BeforeEach(func() {
		setAndUnpausePipeline("fixtures/volume-mounting.yml")
	})

	It("procceds through the plan with input under output mounts", func(ctx SpecContext) {
		watch := fly("trigger-job", "-j", inPipeline("input-under-output"), "-w")
		Expect(watch).To(gexec.Exit(0))
		Expect(watch).To(gbytes.Say("some-resource"))
	}, DefaultSpecTimeout)

	It("procceds through the plan with input under input mounts", func(ctx SpecContext) {
		sess := fly("trigger-job", "-j", inPipeline("input-under-input"), "-w")
		Expect(sess).To(gexec.Exit(0))
		Expect(sess).To(gbytes.Say("helloworld"))
	}, DefaultSpecTimeout)

	It("procceds through the plan having output being mapped to dot and input within", func(ctx SpecContext) {
		sess := fly("trigger-job", "-j", inPipeline("output-with-dot-with-input-within"), "-w")
		Expect(sess).To(gexec.Exit(0))
		Expect(sess).To(gbytes.Say("bar"))
	}, DefaultSpecTimeout)

	It("procceds through the plan with output under input mounts", func(ctx SpecContext) {
		sess := fly("trigger-job", "-j", inPipeline("output-under-input"), "-w")
		Expect(sess).To(gexec.Exit(0))
		Expect(sess).To(gbytes.Say("hello"))
	}, DefaultSpecTimeout)

	It("procceds through the plan with input same as output mounts", func(ctx SpecContext) {
		sess := fly("trigger-job", "-j", inPipeline("input-same-output"), "-w")
		Expect(sess).To(gexec.Exit(0))
		Expect(sess).To(gbytes.Say("hello"))
	}, DefaultSpecTimeout)

	It("procceds through the plan with input and output having the same path but a different name", func(ctx SpecContext) {
		sess := fly("trigger-job", "-j", inPipeline("input-output-same-path-diff-name"), "-w")
		Expect(sess).To(gexec.Exit(0))
	}, DefaultSpecTimeout)
})
