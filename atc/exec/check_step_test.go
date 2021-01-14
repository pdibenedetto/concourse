package exec_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/db/dbfakes"
	"github.com/concourse/concourse/atc/db/lock/lockfakes"
	"github.com/concourse/concourse/atc/exec"
	"github.com/concourse/concourse/atc/exec/execfakes"
	"github.com/concourse/concourse/atc/resource"
	"github.com/concourse/concourse/atc/resource/resourcefakes"
	"github.com/concourse/concourse/atc/runtime"
	"github.com/concourse/concourse/atc/worker"
	"github.com/concourse/concourse/atc/worker/workerfakes"
	"github.com/concourse/concourse/tracing"
	"github.com/concourse/concourse/vars"
	"go.opentelemetry.io/otel/oteltest"
	"go.opentelemetry.io/otel/trace"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CheckStep", func() {
	var (
		ctx    context.Context
		cancel context.CancelFunc

		planID                    atc.PlanID
		fakeRunState              *execfakes.FakeRunState
		fakeResourceFactory       *resourcefakes.FakeResourceFactory
		fakeResource              *resourcefakes.FakeResource
		fakeResourceConfigFactory *dbfakes.FakeResourceConfigFactory
		fakeResourceConfig        *dbfakes.FakeResourceConfig
		fakeResourceConfigScope   *dbfakes.FakeResourceConfigScope
		fakePool                  *workerfakes.FakePool
		fakeClient                *workerfakes.FakeClient
		fakeStrategy              *workerfakes.FakeContainerPlacementStrategy
		fakeDelegate              *execfakes.FakeCheckDelegate
		fakeDelegateFactory       *execfakes.FakeCheckDelegateFactory
		spanCtx                   context.Context
		defaultTimeout            = time.Hour

		fakeStdout, fakeStderr io.Writer

		stepMetadata      exec.StepMetadata
		checkStep         exec.Step
		checkPlan         atc.CheckPlan
		containerMetadata db.ContainerMetadata

		stepOk  bool
		stepErr error
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())

		planID = "some-plan-id"

		fakeRunState = new(execfakes.FakeRunState)
		fakeResourceFactory = new(resourcefakes.FakeResourceFactory)
		fakeResource = new(resourcefakes.FakeResource)
		fakeStrategy = new(workerfakes.FakeContainerPlacementStrategy)
		fakeDelegateFactory = new(execfakes.FakeCheckDelegateFactory)
		fakeDelegate = new(execfakes.FakeCheckDelegate)

		fakeClient = new(workerfakes.FakeClient)
		fakeClient.NameReturns("some-worker")
		fakePool = new(workerfakes.FakePool)
		fakePool.SelectWorkerReturns(fakeClient, 0, nil)

		spanCtx = context.Background()
		fakeDelegate.StartSpanReturns(spanCtx, tracing.NoopSpan)

		fakeStdout = bytes.NewBufferString("out")
		fakeDelegate.StdoutReturns(fakeStdout)

		fakeStderr = bytes.NewBufferString("err")
		fakeDelegate.StderrReturns(fakeStderr)

		stepMetadata = exec.StepMetadata{}
		containerMetadata = db.ContainerMetadata{}

		fakeResourceFactory.NewResourceReturns(fakeResource)

		fakeResourceConfigFactory = new(dbfakes.FakeResourceConfigFactory)
		fakeResourceConfig = new(dbfakes.FakeResourceConfig)
		fakeResourceConfig.IDReturns(501)
		fakeResourceConfig.OriginBaseResourceTypeReturns(&db.UsedBaseResourceType{
			ID:   502,
			Name: "some-base-type",
		})
		fakeResourceConfigFactory.FindOrCreateResourceConfigReturns(fakeResourceConfig, nil)

		fakeResourceConfigScope = new(dbfakes.FakeResourceConfigScope)
		fakeDelegate.FindOrCreateScopeReturns(fakeResourceConfigScope, nil)

		fakeDelegateFactory.CheckDelegateReturns(fakeDelegate)

		checkPlan = atc.CheckPlan{
			Name:    "some-name",
			Type:    "some-base-type",
			Source:  atc.Source{"some": "((source-var))"},
			Timeout: "10s",
			TypeImage: atc.TypeImage{
				BaseType: "some-base-type",
			},
		}

		containerMetadata = db.ContainerMetadata{
			User: "test-user",
		}

		stepMetadata = exec.StepMetadata{
			TeamID:  345,
			BuildID: 678,
		}

		fakeDelegate.VariablesReturns(vars.StaticVariables{"source-var": "super-secret-source"})
	})

	AfterEach(func() {
		cancel()
	})

	JustBeforeEach(func() {
		checkStep = exec.NewCheckStep(
			planID,
			checkPlan,
			stepMetadata,
			fakeResourceFactory,
			fakeResourceConfigFactory,
			containerMetadata,
			fakeStrategy,
			fakePool,
			fakeDelegateFactory,
			defaultTimeout,
		)

		stepOk, stepErr = checkStep.Run(ctx, fakeRunState)
	})

	Context("with a reasonable configuration", func() {
		It("emits an Initializing event", func() {
			Expect(fakeDelegate.InitializingCallCount()).To(Equal(1))
		})

		Context("when not running", func() {
			BeforeEach(func() {
				fakeDelegate.WaitToRunReturns(nil, false, nil)
			})

			It("does not run the check step", func() {
				Expect(fakeClient.RunCheckStepCallCount()).To(Equal(0))
			})

			It("succeeds", func() {
				Expect(stepOk).To(BeTrue())
			})

			Context("when there is a latest version", func() {
				BeforeEach(func() {
					fakeVersion := new(dbfakes.FakeResourceConfigVersion)
					fakeVersion.VersionReturns(db.Version{"some": "latest-version"})
					fakeResourceConfigScope.LatestVersionReturns(fakeVersion, true, nil)
				})

				It("stores the latest version as the step result", func() {
					Expect(fakeRunState.StoreResultCallCount()).To(Equal(1))
					id, val := fakeRunState.StoreResultArgsForCall(0)
					Expect(id).To(Equal(atc.PlanID("some-plan-id")))
					Expect(val).To(Equal(atc.Version{"some": "latest-version"}))
				})
			})

			Context("when there is no version", func() {
				BeforeEach(func() {
					fakeResourceConfigScope.LatestVersionReturns(nil, false, nil)
				})

				It("does not store a version", func() {
					Expect(fakeRunState.StoreResultCallCount()).To(Equal(0))
				})
			})
		})

		Context("running", func() {
			var fakeLock *lockfakes.FakeLock

			BeforeEach(func() {
				fakeLock = new(lockfakes.FakeLock)
				fakeDelegate.WaitToRunReturns(fakeLock, true, nil)
			})

			Context("when given a from version", func() {
				BeforeEach(func() {
					checkPlan.FromVersion = atc.Version{"from": "version"}
				})

				It("constructs the resource with the version", func() {
					Expect(fakeResourceFactory.NewResourceCallCount()).To(Equal(1))
					_, _, fromVersion := fakeResourceFactory.NewResourceArgsForCall(0)
					Expect(fromVersion).To(Equal(checkPlan.FromVersion))
				})
			})

			Context("when not given a from version", func() {
				var fakeVersion *dbfakes.FakeResourceConfigVersion

				BeforeEach(func() {
					checkPlan.FromVersion = nil

					fakeVersion = new(dbfakes.FakeResourceConfigVersion)
					fakeVersion.VersionReturns(db.Version{"latest": "version"})
					fakeResourceConfigScope.LatestVersionStub = func() (db.ResourceConfigVersion, bool, error) {
						Expect(fakeDelegate.WaitToRunCallCount()).To(
							Equal(1),
							"should have gotten latest version after waiting, not before",
						)

						return fakeVersion, true, nil
					}
				})

				It("finds the latest version itself - it's a strong, independent check step who dont need no plan", func() {
					Expect(fakeResourceFactory.NewResourceCallCount()).To(Equal(1))
					_, _, fromVersion := fakeResourceFactory.NewResourceArgsForCall(0)
					Expect(fromVersion).To(Equal(atc.Version{"latest": "version"}))
				})
			})

			Describe("worker selection", func() {
				var ctx context.Context
				var workerSpec worker.WorkerSpec

				JustBeforeEach(func() {
					Expect(fakePool.SelectWorkerCallCount()).To(Equal(1))
					ctx, _, _, workerSpec, _, _ = fakePool.SelectWorkerArgsForCall(0)
				})

				It("doesn't enforce a timeout", func() {
					_, ok := ctx.Deadline()
					Expect(ok).To(BeFalse())
				})

				Describe("calls SelectWorker with the correct WorkerSpec", func() {
					It("with resource type", func() {
						Expect(workerSpec.ResourceType).To(Equal("some-base-type"))
					})

					It("with teamid", func() {
						Expect(workerSpec.TeamID).To(Equal(345))
					})

					Context("when the plan specifies tags", func() {
						BeforeEach(func() {
							checkPlan.Tags = atc.Tags{"some", "tags"}
						})

						It("sets them in the WorkerSpec", func() {
							Expect(workerSpec.Tags).To(Equal([]string{"some", "tags"}))
						})
					})
				})

				It("emits a SelectedWorker event", func() {
					Expect(fakeDelegate.SelectedWorkerCallCount()).To(Equal(1))
					_, workerName := fakeDelegate.SelectedWorkerArgsForCall(0)
					Expect(workerName).To(Equal("some-worker"))
				})

				Context("when selecting a worker fails", func() {
					BeforeEach(func() {
						fakePool.SelectWorkerReturns(nil, 0, errors.New("nope"))
					})

					It("returns an err", func() {
						Expect(stepErr).To(MatchError(ContainSubstring("nope")))
					})
				})
			})

			Describe("running the check step", func() {
				var runCtx context.Context
				var owner db.ContainerOwner
				var containerSpec worker.ContainerSpec
				var metadata db.ContainerMetadata
				var processSpec runtime.ProcessSpec
				var startEventDelegate runtime.StartingEventDelegate
				var resource resource.Resource

				JustBeforeEach(func() {
					Expect(fakeClient.RunCheckStepCallCount()).To(Equal(1), "check step should have run")
					runCtx, owner, containerSpec, metadata, processSpec, startEventDelegate, resource = fakeClient.RunCheckStepArgsForCall(0)
				})

				It("uses ResourceConfigCheckSessionOwner", func() {
					expected := db.NewBuildStepContainerOwner(
						678,
						planID,
						345,
					)

					Expect(owner).To(Equal(expected))
				})

				Context("when using a custom resource type", func() {
					var (
						fakeImageSpec          worker.ImageSpec
						fakeImageResourceCache *dbfakes.FakeResourceCache
					)

					BeforeEach(func() {
						checkPlan.TypeImage.GetPlan = &atc.Plan{
							ID: "1/image-get",
							Get: &atc.GetPlan{
								Name:   "some-custom-type",
								Type:   "another-custom-type",
								Source: atc.Source{"some-custom": "((source-var))"},
								Params: atc.Params{"some-custom": "((params-var))"},
							},
						}

						checkPlan.TypeImage.CheckPlan = &atc.Plan{
							ID: "1/image-check",
							Check: &atc.CheckPlan{
								Name:   "some-custom-type",
								Type:   "another-custom-type",
								Source: atc.Source{"some-custom": "((source-var))"},
							},
						}

						checkPlan.Type = "some-custom-type"

						fakeImageSpec = worker.ImageSpec{
							ImageArtifactSource: new(workerfakes.FakeStreamableArtifactSource),
						}

						fakeImageResourceCache = new(dbfakes.FakeResourceCache)
						fakeImageResourceCache.IDReturns(123)

						fakeDelegate.FetchImageReturns(fakeImageSpec, fakeImageResourceCache, nil)
					})

					It("fetches the resource type image", func() {
						Expect(fakeDelegate.FetchImageCallCount()).To(Equal(1))
						_, actualGetImagePlan, actualCheckImagePlan, privileged := fakeDelegate.FetchImageArgsForCall(0)
						Expect(actualGetImagePlan).To(Equal(*checkPlan.TypeImage.GetPlan))
						Expect(actualCheckImagePlan).To(Equal(checkPlan.TypeImage.CheckPlan))
						Expect(privileged).To(BeFalse())
					})

					It("sets the image spec in the container spec", func() {
						Expect(containerSpec.ImageSpec).To(Equal(fakeImageSpec))
					})

					It("creates the resource config using the image resource cache", func() {
						Expect(fakeResourceConfigFactory.FindOrCreateResourceConfigCallCount()).To(Equal(1))
						type_, source, irc := fakeResourceConfigFactory.FindOrCreateResourceConfigArgsForCall(0)
						Expect(type_).To(Equal("some-custom-type"))
						Expect(source).To(Equal(atc.Source{"some": "super-secret-source"}))
						Expect(irc).To(Equal(fakeImageResourceCache))
					})

					Context("when the resource type is privileged", func() {
						BeforeEach(func() {
							checkPlan.TypeImage.Privileged = true
						})

						It("fetches the image with privileged", func() {
							Expect(fakeDelegate.FetchImageCallCount()).To(Equal(1))
							_, _, _, privileged := fakeDelegate.FetchImageArgsForCall(0)
							Expect(privileged).To(BeTrue())
						})
					})
				})

				Context("when the plan is for a resource", func() {
					BeforeEach(func() {
						checkPlan.Resource = "some-resource"
					})

					It("uses ResourceConfigCheckSessionOwner", func() {
						expected := db.NewResourceConfigCheckSessionContainerOwner(
							501,
							502,
							db.ContainerOwnerExpiries{Min: 5 * time.Minute, Max: 1 * time.Hour},
						)

						Expect(owner).To(Equal(expected))
					})
				})

				Context("when the plan specifies a timeout", func() {
					BeforeEach(func() {
						checkPlan.Timeout = "1h"
					})

					It("enforces it on the check", func() {
						t, ok := runCtx.Deadline()
						Expect(ok).To(BeTrue())
						Expect(t).To(BeTemporally("~", time.Now().Add(time.Hour), time.Minute))
					})

					Context("when running times out", func() {
						BeforeEach(func() {
							fakeClient.RunCheckStepReturns(
								worker.CheckResult{},
								fmt.Errorf("wrapped: %w", context.DeadlineExceeded),
							)
						})

						It("fails without error", func() {
							Expect(stepOk).To(BeFalse())
							Expect(stepErr).To(BeNil())
						})

						It("emits an Errored event", func() {
							Expect(fakeDelegate.ErroredCallCount()).To(Equal(1))
							_, status := fakeDelegate.ErroredArgsForCall(0)
							Expect(status).To(Equal(exec.TimeoutLogMessage))
						})
					})
				})

				It("passes the process spec", func() {
					Expect(processSpec).To(Equal(runtime.ProcessSpec{
						Path:         "/opt/resource/check",
						StdoutWriter: fakeStdout,
						StderrWriter: fakeStderr,
					}))
				})

				It("passes the delegate as the start event delegate", func() {
					Expect(startEventDelegate).To(Equal(fakeDelegate))
				})

				Context("uses containerspec", func() {
					It("with certs volume mount", func() {
						Expect(containerSpec.BindMounts).To(HaveLen(1))
						mount := containerSpec.BindMounts[0]

						_, ok := mount.(*worker.CertsVolumeMount)
						Expect(ok).To(BeTrue())
					})

					It("uses base type for image", func() {
						Expect(containerSpec.ImageSpec).To(Equal(worker.ImageSpec{
							ResourceType: "some-base-type",
						}))
					})

					It("with teamid set", func() {
						Expect(containerSpec.TeamID).To(Equal(345))
					})

					It("with env vars", func() {
						Expect(containerSpec.Env).To(ContainElement("BUILD_TEAM_ID=345"))
					})

					Context("when tracing is enabled", func() {
						var buildSpan trace.Span

						BeforeEach(func() {
							tracing.ConfigureTraceProvider(oteltest.NewTracerProvider())

							spanCtx, buildSpan = tracing.StartSpan(ctx, "build", nil)
							fakeDelegate.StartSpanReturns(spanCtx, buildSpan)
						})

						AfterEach(func() {
							tracing.Configured = false
						})

						It("propagates span context to the worker client", func() {
							Expect(runCtx).To(Equal(rewrapLogger(spanCtx)))
						})

						It("populates the TRACEPARENT env var", func() {
							Expect(containerSpec.Env).To(ContainElement(MatchRegexp(`TRACEPARENT=.+`)))
						})
					})
				})

				It("uses container metadata", func() {
					Expect(metadata).To(Equal(containerMetadata))
				})

				It("uses the resource created", func() {
					Expect(resource).To(Equal(fakeResource))
				})
			})

			Context("with tracing configured", func() {
				var buildSpan trace.Span

				BeforeEach(func() {
					tracing.ConfigureTraceProvider(oteltest.NewTracerProvider())

					spanCtx, buildSpan = tracing.StartSpan(context.Background(), "fake-operation", nil)
					fakeDelegate.StartSpanReturns(spanCtx, buildSpan)
				})

				AfterEach(func() {
					tracing.Configured = false
				})

				It("propagates span context to scope", func() {
					Expect(fakeResourceConfigScope.SaveVersionsCallCount()).To(Equal(1))
					spanContext, _ := fakeResourceConfigScope.SaveVersionsArgsForCall(0)
					traceID := buildSpan.SpanContext().TraceID().String()
					traceParent := spanContext.Get("traceparent")
					Expect(traceParent).To(ContainSubstring(traceID))
				})
			})

			Context("having RunCheckStep succeed", func() {
				BeforeEach(func() {
					fakeClient.RunCheckStepReturns(worker.CheckResult{
						Versions: []atc.Version{
							{"version": "1"},
							{"version": "2"},
						},
					}, nil)
				})

				It("succeeds", func() {
					Expect(stepOk).To(BeTrue())
				})

				It("saves the versions to the config scope", func() {
					Expect(fakeResourceConfigFactory.FindOrCreateResourceConfigCallCount()).To(Equal(1))
					type_, source, irc := fakeResourceConfigFactory.FindOrCreateResourceConfigArgsForCall(0)
					Expect(type_).To(Equal("some-base-type"))
					Expect(source).To(Equal(atc.Source{"some": "super-secret-source"}))
					Expect(irc).To(BeNil())

					Expect(fakeDelegate.FindOrCreateScopeCallCount()).To(Equal(1))
					config := fakeDelegate.FindOrCreateScopeArgsForCall(0)
					Expect(config).To(Equal(fakeResourceConfig))

					spanContext, versions := fakeResourceConfigScope.SaveVersionsArgsForCall(0)
					Expect(spanContext).To(Equal(db.SpanContext{}))
					Expect(versions).To(Equal([]atc.Version{
						{"version": "1"},
						{"version": "2"},
					}))
				})

				It("stores the latest version as the step result", func() {
					Expect(fakeRunState.StoreResultCallCount()).To(Equal(1))
					id, val := fakeRunState.StoreResultArgsForCall(0)
					Expect(id).To(Equal(atc.PlanID("some-plan-id")))
					Expect(val).To(Equal(atc.Version{"version": "2"}))
				})

				It("emits a successful Finished event", func() {
					Expect(fakeDelegate.FinishedCallCount()).To(Equal(1))
					_, succeeded := fakeDelegate.FinishedArgsForCall(0)
					Expect(succeeded).To(BeTrue())
				})

				Context("when no versions are returned", func() {
					BeforeEach(func() {
						fakeClient.RunCheckStepReturns(worker.CheckResult{Versions: []atc.Version{}}, nil)
					})

					It("succeeds", func() {
						Expect(stepErr).ToNot(HaveOccurred())
						Expect(stepOk).To(BeTrue())
					})

					It("does not store a version", func() {
						Expect(fakeRunState.StoreResultCallCount()).To(Equal(0))
					})
				})

				Context("before running the check", func() {
					BeforeEach(func() {
						fakeResourceConfigScope.UpdateLastCheckStartTimeStub = func() (bool, error) {
							Expect(fakeClient.RunCheckStepCallCount()).To(Equal(0))
							return true, nil
						}
					})

					It("updates the scope's last check start time", func() {
						Expect(fakeResourceConfigScope.UpdateLastCheckStartTimeCallCount()).To(Equal(1))
						Expect(fakeClient.RunCheckStepCallCount()).To(Equal(1))
					})
				})

				Context("after saving", func() {
					BeforeEach(func() {
						fakeResourceConfigScope.SaveVersionsStub = func(db.SpanContext, []atc.Version) error {
							Expect(fakeDelegate.PointToCheckedConfigCallCount()).To(BeZero())
							Expect(fakeResourceConfigScope.UpdateLastCheckEndTimeCallCount()).To(Equal(0))
							return nil
						}
					})

					It("updates the scope's last check end time", func() {
						Expect(fakeResourceConfigScope.UpdateLastCheckEndTimeCallCount()).To(Equal(1))
					})

					It("points the resource or resource type to the scope", func() {
						Expect(fakeResourceConfigScope.SaveVersionsCallCount()).To(Equal(1))
						Expect(fakeDelegate.PointToCheckedConfigCallCount()).To(Equal(1))
						scope := fakeDelegate.PointToCheckedConfigArgsForCall(0)
						Expect(scope).To(Equal(fakeResourceConfigScope))
					})
				})

				Context("after pointing the resource type to the scope", func() {
					BeforeEach(func() {
						fakeDelegate.PointToCheckedConfigStub = func(db.ResourceConfigScope) error {
							Expect(fakeLock.ReleaseCallCount()).To(Equal(0))
							return nil
						}
					})

					It("releases the lock", func() {
						Expect(fakeDelegate.PointToCheckedConfigCallCount()).To(Equal(1))
						Expect(fakeLock.ReleaseCallCount()).To(Equal(1))
					})
				})
			})

			Context("having RunCheckStep erroring", func() {
				var expectedErr error

				BeforeEach(func() {
					expectedErr = errors.New("run-check-step-err")
					fakeClient.RunCheckStepReturns(worker.CheckResult{}, expectedErr)
				})

				It("errors", func() {
					Expect(stepErr).To(HaveOccurred())
					Expect(errors.Is(stepErr, expectedErr)).To(BeTrue())
				})

				It("points the resource or resource type to the scope", func() {
					// even though we failed to check, we should still point to the new
					// scope; it'd be kind of weird leave the resource pointing to the old
					// scope for a substantial config change that also happens to be
					// broken.
					Expect(fakeDelegate.PointToCheckedConfigCallCount()).To(Equal(1))
					scope := fakeDelegate.PointToCheckedConfigArgsForCall(0)
					Expect(scope).To(Equal(fakeResourceConfigScope))
				})

				It("updates the scope's last check end time", func() {
					Expect(fakeResourceConfigScope.UpdateLastCheckEndTimeCallCount()).To(Equal(1))
				})

				// Finished is for script success/failure, whereas this is an error
				It("does not emit a Finished event", func() {
					Expect(fakeDelegate.FinishedCallCount()).To(Equal(0))
				})

				Context("with a script failure", func() {
					BeforeEach(func() {
						fakeClient.RunCheckStepReturns(worker.CheckResult{}, runtime.ErrResourceScriptFailed{
							ExitStatus: 42,
						})
					})

					It("does not error", func() {
						// don't return an error - the script output has already been
						// printed, and emitting an errored event would double it up
						Expect(stepErr).ToNot(HaveOccurred())
					})

					It("updates the scope's last check end time", func() {
						Expect(fakeResourceConfigScope.UpdateLastCheckEndTimeCallCount()).To(Equal(1))
					})

					It("emits a failed Finished event", func() {
						Expect(fakeDelegate.FinishedCallCount()).To(Equal(1))
						_, succeeded := fakeDelegate.FinishedArgsForCall(0)
						Expect(succeeded).To(BeFalse())
					})
				})
			})

			Context("having SaveVersions failing", func() {
				var expectedErr error

				BeforeEach(func() {
					expectedErr = errors.New("save-versions-err")

					fakeResourceConfigScope.SaveVersionsReturns(expectedErr)
				})

				It("errors", func() {
					Expect(stepErr).To(HaveOccurred())
					Expect(errors.Is(stepErr, expectedErr)).To(BeTrue())
				})
			})
		})
	})

	Context("having credentials in the config", func() {
		BeforeEach(func() {
			checkPlan.Source = atc.Source{"some": "((super-secret-source))"}
		})

		Context("having cred evaluation failing", func() {
			var expectedErr error

			BeforeEach(func() {
				expectedErr = errors.New("creds-err")
			})

			It("errors", func() {
				Expect(stepErr).To(HaveOccurred())
				Expect(errors.Is(stepErr, expectedErr)).To(BeTrue())
			})
		})
	})

	Context("when a bogus timeout is given", func() {
		BeforeEach(func() {
			checkPlan.Timeout = "bogus"
		})

		It("fails miserably", func() {
			Expect(stepErr).To(MatchError("parse timeout: time: invalid duration \"bogus\""))
		})
	})
})
