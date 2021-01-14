// Code generated by counterfeiter. DO NOT EDIT.
package execfakes

import (
	"context"
	"io"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/exec"
	"github.com/concourse/concourse/atc/runtime"
	"github.com/concourse/concourse/atc/worker"
	"github.com/concourse/concourse/tracing"
	"github.com/concourse/concourse/vars"
	"go.opentelemetry.io/otel/trace"
)

type FakePutDelegate struct {
	ErroredStub        func(lager.Logger, string)
	erroredMutex       sync.RWMutex
	erroredArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	FetchImageStub        func(context.Context, atc.Plan, *atc.Plan, bool) (worker.ImageSpec, db.ResourceCache, error)
	fetchImageMutex       sync.RWMutex
	fetchImageArgsForCall []struct {
		arg1 context.Context
		arg2 atc.Plan
		arg3 *atc.Plan
		arg4 bool
	}
	fetchImageReturns struct {
		result1 worker.ImageSpec
		result2 db.ResourceCache
		result3 error
	}
	fetchImageReturnsOnCall map[int]struct {
		result1 worker.ImageSpec
		result2 db.ResourceCache
		result3 error
	}
	FinishedStub        func(lager.Logger, exec.ExitStatus, runtime.VersionResult)
	finishedMutex       sync.RWMutex
	finishedArgsForCall []struct {
		arg1 lager.Logger
		arg2 exec.ExitStatus
		arg3 runtime.VersionResult
	}
	InitializingStub        func(lager.Logger)
	initializingMutex       sync.RWMutex
	initializingArgsForCall []struct {
		arg1 lager.Logger
	}
	SaveOutputStub        func(lager.Logger, atc.PutPlan, atc.Source, db.ResourceCache, runtime.VersionResult)
	saveOutputMutex       sync.RWMutex
	saveOutputArgsForCall []struct {
		arg1 lager.Logger
		arg2 atc.PutPlan
		arg3 atc.Source
		arg4 db.ResourceCache
		arg5 runtime.VersionResult
	}
	SelectedWorkerStub        func(lager.Logger, string)
	selectedWorkerMutex       sync.RWMutex
	selectedWorkerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	StartSpanStub        func(context.Context, string, tracing.Attrs) (context.Context, trace.Span)
	startSpanMutex       sync.RWMutex
	startSpanArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 tracing.Attrs
	}
	startSpanReturns struct {
		result1 context.Context
		result2 trace.Span
	}
	startSpanReturnsOnCall map[int]struct {
		result1 context.Context
		result2 trace.Span
	}
	StartingStub        func(lager.Logger)
	startingMutex       sync.RWMutex
	startingArgsForCall []struct {
		arg1 lager.Logger
	}
	StderrStub        func() io.Writer
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct {
	}
	stderrReturns struct {
		result1 io.Writer
	}
	stderrReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	StdoutStub        func() io.Writer
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct {
	}
	stdoutReturns struct {
		result1 io.Writer
	}
	stdoutReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	VariablesStub        func(context.Context) vars.Variables
	variablesMutex       sync.RWMutex
	variablesArgsForCall []struct {
		arg1 context.Context
	}
	variablesReturns struct {
		result1 vars.Variables
	}
	variablesReturnsOnCall map[int]struct {
		result1 vars.Variables
	}
	WaitingForWorkerStub        func(lager.Logger)
	waitingForWorkerMutex       sync.RWMutex
	waitingForWorkerArgsForCall []struct {
		arg1 lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePutDelegate) Errored(arg1 lager.Logger, arg2 string) {
	fake.erroredMutex.Lock()
	fake.erroredArgsForCall = append(fake.erroredArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.ErroredStub
	fake.recordInvocation("Errored", []interface{}{arg1, arg2})
	fake.erroredMutex.Unlock()
	if stub != nil {
		fake.ErroredStub(arg1, arg2)
	}
}

func (fake *FakePutDelegate) ErroredCallCount() int {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	return len(fake.erroredArgsForCall)
}

func (fake *FakePutDelegate) ErroredCalls(stub func(lager.Logger, string)) {
	fake.erroredMutex.Lock()
	defer fake.erroredMutex.Unlock()
	fake.ErroredStub = stub
}

func (fake *FakePutDelegate) ErroredArgsForCall(i int) (lager.Logger, string) {
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	argsForCall := fake.erroredArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePutDelegate) FetchImage(arg1 context.Context, arg2 atc.Plan, arg3 *atc.Plan, arg4 bool) (worker.ImageSpec, db.ResourceCache, error) {
	fake.fetchImageMutex.Lock()
	ret, specificReturn := fake.fetchImageReturnsOnCall[len(fake.fetchImageArgsForCall)]
	fake.fetchImageArgsForCall = append(fake.fetchImageArgsForCall, struct {
		arg1 context.Context
		arg2 atc.Plan
		arg3 *atc.Plan
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	stub := fake.FetchImageStub
	fakeReturns := fake.fetchImageReturns
	fake.recordInvocation("FetchImage", []interface{}{arg1, arg2, arg3, arg4})
	fake.fetchImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakePutDelegate) FetchImageCallCount() int {
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	return len(fake.fetchImageArgsForCall)
}

func (fake *FakePutDelegate) FetchImageCalls(stub func(context.Context, atc.Plan, *atc.Plan, bool) (worker.ImageSpec, db.ResourceCache, error)) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = stub
}

func (fake *FakePutDelegate) FetchImageArgsForCall(i int) (context.Context, atc.Plan, *atc.Plan, bool) {
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	argsForCall := fake.fetchImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakePutDelegate) FetchImageReturns(result1 worker.ImageSpec, result2 db.ResourceCache, result3 error) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = nil
	fake.fetchImageReturns = struct {
		result1 worker.ImageSpec
		result2 db.ResourceCache
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePutDelegate) FetchImageReturnsOnCall(i int, result1 worker.ImageSpec, result2 db.ResourceCache, result3 error) {
	fake.fetchImageMutex.Lock()
	defer fake.fetchImageMutex.Unlock()
	fake.FetchImageStub = nil
	if fake.fetchImageReturnsOnCall == nil {
		fake.fetchImageReturnsOnCall = make(map[int]struct {
			result1 worker.ImageSpec
			result2 db.ResourceCache
			result3 error
		})
	}
	fake.fetchImageReturnsOnCall[i] = struct {
		result1 worker.ImageSpec
		result2 db.ResourceCache
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePutDelegate) Finished(arg1 lager.Logger, arg2 exec.ExitStatus, arg3 runtime.VersionResult) {
	fake.finishedMutex.Lock()
	fake.finishedArgsForCall = append(fake.finishedArgsForCall, struct {
		arg1 lager.Logger
		arg2 exec.ExitStatus
		arg3 runtime.VersionResult
	}{arg1, arg2, arg3})
	stub := fake.FinishedStub
	fake.recordInvocation("Finished", []interface{}{arg1, arg2, arg3})
	fake.finishedMutex.Unlock()
	if stub != nil {
		fake.FinishedStub(arg1, arg2, arg3)
	}
}

func (fake *FakePutDelegate) FinishedCallCount() int {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	return len(fake.finishedArgsForCall)
}

func (fake *FakePutDelegate) FinishedCalls(stub func(lager.Logger, exec.ExitStatus, runtime.VersionResult)) {
	fake.finishedMutex.Lock()
	defer fake.finishedMutex.Unlock()
	fake.FinishedStub = stub
}

func (fake *FakePutDelegate) FinishedArgsForCall(i int) (lager.Logger, exec.ExitStatus, runtime.VersionResult) {
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	argsForCall := fake.finishedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePutDelegate) Initializing(arg1 lager.Logger) {
	fake.initializingMutex.Lock()
	fake.initializingArgsForCall = append(fake.initializingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.InitializingStub
	fake.recordInvocation("Initializing", []interface{}{arg1})
	fake.initializingMutex.Unlock()
	if stub != nil {
		fake.InitializingStub(arg1)
	}
}

func (fake *FakePutDelegate) InitializingCallCount() int {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	return len(fake.initializingArgsForCall)
}

func (fake *FakePutDelegate) InitializingCalls(stub func(lager.Logger)) {
	fake.initializingMutex.Lock()
	defer fake.initializingMutex.Unlock()
	fake.InitializingStub = stub
}

func (fake *FakePutDelegate) InitializingArgsForCall(i int) lager.Logger {
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	argsForCall := fake.initializingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePutDelegate) SaveOutput(arg1 lager.Logger, arg2 atc.PutPlan, arg3 atc.Source, arg4 db.ResourceCache, arg5 runtime.VersionResult) {
	fake.saveOutputMutex.Lock()
	fake.saveOutputArgsForCall = append(fake.saveOutputArgsForCall, struct {
		arg1 lager.Logger
		arg2 atc.PutPlan
		arg3 atc.Source
		arg4 db.ResourceCache
		arg5 runtime.VersionResult
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.SaveOutputStub
	fake.recordInvocation("SaveOutput", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.saveOutputMutex.Unlock()
	if stub != nil {
		fake.SaveOutputStub(arg1, arg2, arg3, arg4, arg5)
	}
}

func (fake *FakePutDelegate) SaveOutputCallCount() int {
	fake.saveOutputMutex.RLock()
	defer fake.saveOutputMutex.RUnlock()
	return len(fake.saveOutputArgsForCall)
}

func (fake *FakePutDelegate) SaveOutputCalls(stub func(lager.Logger, atc.PutPlan, atc.Source, db.ResourceCache, runtime.VersionResult)) {
	fake.saveOutputMutex.Lock()
	defer fake.saveOutputMutex.Unlock()
	fake.SaveOutputStub = stub
}

func (fake *FakePutDelegate) SaveOutputArgsForCall(i int) (lager.Logger, atc.PutPlan, atc.Source, db.ResourceCache, runtime.VersionResult) {
	fake.saveOutputMutex.RLock()
	defer fake.saveOutputMutex.RUnlock()
	argsForCall := fake.saveOutputArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakePutDelegate) SelectedWorker(arg1 lager.Logger, arg2 string) {
	fake.selectedWorkerMutex.Lock()
	fake.selectedWorkerArgsForCall = append(fake.selectedWorkerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	stub := fake.SelectedWorkerStub
	fake.recordInvocation("SelectedWorker", []interface{}{arg1, arg2})
	fake.selectedWorkerMutex.Unlock()
	if stub != nil {
		fake.SelectedWorkerStub(arg1, arg2)
	}
}

func (fake *FakePutDelegate) SelectedWorkerCallCount() int {
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	return len(fake.selectedWorkerArgsForCall)
}

func (fake *FakePutDelegate) SelectedWorkerCalls(stub func(lager.Logger, string)) {
	fake.selectedWorkerMutex.Lock()
	defer fake.selectedWorkerMutex.Unlock()
	fake.SelectedWorkerStub = stub
}

func (fake *FakePutDelegate) SelectedWorkerArgsForCall(i int) (lager.Logger, string) {
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	argsForCall := fake.selectedWorkerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePutDelegate) StartSpan(arg1 context.Context, arg2 string, arg3 tracing.Attrs) (context.Context, trace.Span) {
	fake.startSpanMutex.Lock()
	ret, specificReturn := fake.startSpanReturnsOnCall[len(fake.startSpanArgsForCall)]
	fake.startSpanArgsForCall = append(fake.startSpanArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 tracing.Attrs
	}{arg1, arg2, arg3})
	stub := fake.StartSpanStub
	fakeReturns := fake.startSpanReturns
	fake.recordInvocation("StartSpan", []interface{}{arg1, arg2, arg3})
	fake.startSpanMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePutDelegate) StartSpanCallCount() int {
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	return len(fake.startSpanArgsForCall)
}

func (fake *FakePutDelegate) StartSpanCalls(stub func(context.Context, string, tracing.Attrs) (context.Context, trace.Span)) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = stub
}

func (fake *FakePutDelegate) StartSpanArgsForCall(i int) (context.Context, string, tracing.Attrs) {
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	argsForCall := fake.startSpanArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePutDelegate) StartSpanReturns(result1 context.Context, result2 trace.Span) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = nil
	fake.startSpanReturns = struct {
		result1 context.Context
		result2 trace.Span
	}{result1, result2}
}

func (fake *FakePutDelegate) StartSpanReturnsOnCall(i int, result1 context.Context, result2 trace.Span) {
	fake.startSpanMutex.Lock()
	defer fake.startSpanMutex.Unlock()
	fake.StartSpanStub = nil
	if fake.startSpanReturnsOnCall == nil {
		fake.startSpanReturnsOnCall = make(map[int]struct {
			result1 context.Context
			result2 trace.Span
		})
	}
	fake.startSpanReturnsOnCall[i] = struct {
		result1 context.Context
		result2 trace.Span
	}{result1, result2}
}

func (fake *FakePutDelegate) Starting(arg1 lager.Logger) {
	fake.startingMutex.Lock()
	fake.startingArgsForCall = append(fake.startingArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.StartingStub
	fake.recordInvocation("Starting", []interface{}{arg1})
	fake.startingMutex.Unlock()
	if stub != nil {
		fake.StartingStub(arg1)
	}
}

func (fake *FakePutDelegate) StartingCallCount() int {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	return len(fake.startingArgsForCall)
}

func (fake *FakePutDelegate) StartingCalls(stub func(lager.Logger)) {
	fake.startingMutex.Lock()
	defer fake.startingMutex.Unlock()
	fake.StartingStub = stub
}

func (fake *FakePutDelegate) StartingArgsForCall(i int) lager.Logger {
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	argsForCall := fake.startingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePutDelegate) Stderr() io.Writer {
	fake.stderrMutex.Lock()
	ret, specificReturn := fake.stderrReturnsOnCall[len(fake.stderrArgsForCall)]
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct {
	}{})
	stub := fake.StderrStub
	fakeReturns := fake.stderrReturns
	fake.recordInvocation("Stderr", []interface{}{})
	fake.stderrMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePutDelegate) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakePutDelegate) StderrCalls(stub func() io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = stub
}

func (fake *FakePutDelegate) StderrReturns(result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakePutDelegate) StderrReturnsOnCall(i int, result1 io.Writer) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	if fake.stderrReturnsOnCall == nil {
		fake.stderrReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stderrReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakePutDelegate) Stdout() io.Writer {
	fake.stdoutMutex.Lock()
	ret, specificReturn := fake.stdoutReturnsOnCall[len(fake.stdoutArgsForCall)]
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct {
	}{})
	stub := fake.StdoutStub
	fakeReturns := fake.stdoutReturns
	fake.recordInvocation("Stdout", []interface{}{})
	fake.stdoutMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePutDelegate) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakePutDelegate) StdoutCalls(stub func() io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = stub
}

func (fake *FakePutDelegate) StdoutReturns(result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakePutDelegate) StdoutReturnsOnCall(i int, result1 io.Writer) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	if fake.stdoutReturnsOnCall == nil {
		fake.stdoutReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stdoutReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakePutDelegate) Variables(arg1 context.Context) vars.Variables {
	fake.variablesMutex.Lock()
	ret, specificReturn := fake.variablesReturnsOnCall[len(fake.variablesArgsForCall)]
	fake.variablesArgsForCall = append(fake.variablesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.VariablesStub
	fakeReturns := fake.variablesReturns
	fake.recordInvocation("Variables", []interface{}{arg1})
	fake.variablesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePutDelegate) VariablesCallCount() int {
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	return len(fake.variablesArgsForCall)
}

func (fake *FakePutDelegate) VariablesCalls(stub func(context.Context) vars.Variables) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = stub
}

func (fake *FakePutDelegate) VariablesArgsForCall(i int) context.Context {
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	argsForCall := fake.variablesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePutDelegate) VariablesReturns(result1 vars.Variables) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = nil
	fake.variablesReturns = struct {
		result1 vars.Variables
	}{result1}
}

func (fake *FakePutDelegate) VariablesReturnsOnCall(i int, result1 vars.Variables) {
	fake.variablesMutex.Lock()
	defer fake.variablesMutex.Unlock()
	fake.VariablesStub = nil
	if fake.variablesReturnsOnCall == nil {
		fake.variablesReturnsOnCall = make(map[int]struct {
			result1 vars.Variables
		})
	}
	fake.variablesReturnsOnCall[i] = struct {
		result1 vars.Variables
	}{result1}
}

func (fake *FakePutDelegate) WaitingForWorker(arg1 lager.Logger) {
	fake.waitingForWorkerMutex.Lock()
	fake.waitingForWorkerArgsForCall = append(fake.waitingForWorkerArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	stub := fake.WaitingForWorkerStub
	fake.recordInvocation("WaitingForWorker", []interface{}{arg1})
	fake.waitingForWorkerMutex.Unlock()
	if stub != nil {
		fake.WaitingForWorkerStub(arg1)
	}
}

func (fake *FakePutDelegate) WaitingForWorkerCallCount() int {
	fake.waitingForWorkerMutex.RLock()
	defer fake.waitingForWorkerMutex.RUnlock()
	return len(fake.waitingForWorkerArgsForCall)
}

func (fake *FakePutDelegate) WaitingForWorkerCalls(stub func(lager.Logger)) {
	fake.waitingForWorkerMutex.Lock()
	defer fake.waitingForWorkerMutex.Unlock()
	fake.WaitingForWorkerStub = stub
}

func (fake *FakePutDelegate) WaitingForWorkerArgsForCall(i int) lager.Logger {
	fake.waitingForWorkerMutex.RLock()
	defer fake.waitingForWorkerMutex.RUnlock()
	argsForCall := fake.waitingForWorkerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePutDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.erroredMutex.RLock()
	defer fake.erroredMutex.RUnlock()
	fake.fetchImageMutex.RLock()
	defer fake.fetchImageMutex.RUnlock()
	fake.finishedMutex.RLock()
	defer fake.finishedMutex.RUnlock()
	fake.initializingMutex.RLock()
	defer fake.initializingMutex.RUnlock()
	fake.saveOutputMutex.RLock()
	defer fake.saveOutputMutex.RUnlock()
	fake.selectedWorkerMutex.RLock()
	defer fake.selectedWorkerMutex.RUnlock()
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	fake.startingMutex.RLock()
	defer fake.startingMutex.RUnlock()
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	fake.waitingForWorkerMutex.RLock()
	defer fake.waitingForWorkerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePutDelegate) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ exec.PutDelegate = new(FakePutDelegate)
