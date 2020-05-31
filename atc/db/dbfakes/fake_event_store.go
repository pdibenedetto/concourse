// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/event"
)

type FakeEventStore struct {
	DeleteStub        func(context.Context, []db.Build) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 []db.Build
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeletePipelineStub        func(context.Context, db.Pipeline) error
	deletePipelineMutex       sync.RWMutex
	deletePipelineArgsForCall []struct {
		arg1 context.Context
		arg2 db.Pipeline
	}
	deletePipelineReturns struct {
		result1 error
	}
	deletePipelineReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteTeamStub        func(context.Context, db.Team) error
	deleteTeamMutex       sync.RWMutex
	deleteTeamArgsForCall []struct {
		arg1 context.Context
		arg2 db.Team
	}
	deleteTeamReturns struct {
		result1 error
	}
	deleteTeamReturnsOnCall map[int]struct {
		result1 error
	}
	FinalizeStub        func(context.Context, db.Build) error
	finalizeMutex       sync.RWMutex
	finalizeArgsForCall []struct {
		arg1 context.Context
		arg2 db.Build
	}
	finalizeReturns struct {
		result1 error
	}
	finalizeReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(context.Context, db.Build, int, *db.EventKey) ([]event.Envelope, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 db.Build
		arg3 int
		arg4 *db.EventKey
	}
	getReturns struct {
		result1 []event.Envelope
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 []event.Envelope
		result2 error
	}
	InitializeStub        func(context.Context, db.Build) error
	initializeMutex       sync.RWMutex
	initializeArgsForCall []struct {
		arg1 context.Context
		arg2 db.Build
	}
	initializeReturns struct {
		result1 error
	}
	initializeReturnsOnCall map[int]struct {
		result1 error
	}
	PutStub        func(context.Context, db.Build, []atc.Event) (db.EventKey, error)
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 context.Context
		arg2 db.Build
		arg3 []atc.Event
	}
	putReturns struct {
		result1 db.EventKey
		result2 error
	}
	putReturnsOnCall map[int]struct {
		result1 db.EventKey
		result2 error
	}
	SetupStub        func(context.Context) error
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		arg1 context.Context
	}
	setupReturns struct {
		result1 error
	}
	setupReturnsOnCall map[int]struct {
		result1 error
	}
	UnmarshalKeyStub        func([]byte, *db.EventKey) error
	unmarshalKeyMutex       sync.RWMutex
	unmarshalKeyArgsForCall []struct {
		arg1 []byte
		arg2 *db.EventKey
	}
	unmarshalKeyReturns struct {
		result1 error
	}
	unmarshalKeyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventStore) Delete(arg1 context.Context, arg2 []db.Build) error {
	var arg2Copy []db.Build
	if arg2 != nil {
		arg2Copy = make([]db.Build, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 []db.Build
	}{arg1, arg2Copy})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2Copy})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeEventStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeEventStore) DeleteCalls(stub func(context.Context, []db.Build) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeEventStore) DeleteArgsForCall(i int) (context.Context, []db.Build) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventStore) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) DeletePipeline(arg1 context.Context, arg2 db.Pipeline) error {
	fake.deletePipelineMutex.Lock()
	ret, specificReturn := fake.deletePipelineReturnsOnCall[len(fake.deletePipelineArgsForCall)]
	fake.deletePipelineArgsForCall = append(fake.deletePipelineArgsForCall, struct {
		arg1 context.Context
		arg2 db.Pipeline
	}{arg1, arg2})
	fake.recordInvocation("DeletePipeline", []interface{}{arg1, arg2})
	fake.deletePipelineMutex.Unlock()
	if fake.DeletePipelineStub != nil {
		return fake.DeletePipelineStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deletePipelineReturns
	return fakeReturns.result1
}

func (fake *FakeEventStore) DeletePipelineCallCount() int {
	fake.deletePipelineMutex.RLock()
	defer fake.deletePipelineMutex.RUnlock()
	return len(fake.deletePipelineArgsForCall)
}

func (fake *FakeEventStore) DeletePipelineCalls(stub func(context.Context, db.Pipeline) error) {
	fake.deletePipelineMutex.Lock()
	defer fake.deletePipelineMutex.Unlock()
	fake.DeletePipelineStub = stub
}

func (fake *FakeEventStore) DeletePipelineArgsForCall(i int) (context.Context, db.Pipeline) {
	fake.deletePipelineMutex.RLock()
	defer fake.deletePipelineMutex.RUnlock()
	argsForCall := fake.deletePipelineArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventStore) DeletePipelineReturns(result1 error) {
	fake.deletePipelineMutex.Lock()
	defer fake.deletePipelineMutex.Unlock()
	fake.DeletePipelineStub = nil
	fake.deletePipelineReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) DeletePipelineReturnsOnCall(i int, result1 error) {
	fake.deletePipelineMutex.Lock()
	defer fake.deletePipelineMutex.Unlock()
	fake.DeletePipelineStub = nil
	if fake.deletePipelineReturnsOnCall == nil {
		fake.deletePipelineReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deletePipelineReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) DeleteTeam(arg1 context.Context, arg2 db.Team) error {
	fake.deleteTeamMutex.Lock()
	ret, specificReturn := fake.deleteTeamReturnsOnCall[len(fake.deleteTeamArgsForCall)]
	fake.deleteTeamArgsForCall = append(fake.deleteTeamArgsForCall, struct {
		arg1 context.Context
		arg2 db.Team
	}{arg1, arg2})
	fake.recordInvocation("DeleteTeam", []interface{}{arg1, arg2})
	fake.deleteTeamMutex.Unlock()
	if fake.DeleteTeamStub != nil {
		return fake.DeleteTeamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteTeamReturns
	return fakeReturns.result1
}

func (fake *FakeEventStore) DeleteTeamCallCount() int {
	fake.deleteTeamMutex.RLock()
	defer fake.deleteTeamMutex.RUnlock()
	return len(fake.deleteTeamArgsForCall)
}

func (fake *FakeEventStore) DeleteTeamCalls(stub func(context.Context, db.Team) error) {
	fake.deleteTeamMutex.Lock()
	defer fake.deleteTeamMutex.Unlock()
	fake.DeleteTeamStub = stub
}

func (fake *FakeEventStore) DeleteTeamArgsForCall(i int) (context.Context, db.Team) {
	fake.deleteTeamMutex.RLock()
	defer fake.deleteTeamMutex.RUnlock()
	argsForCall := fake.deleteTeamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventStore) DeleteTeamReturns(result1 error) {
	fake.deleteTeamMutex.Lock()
	defer fake.deleteTeamMutex.Unlock()
	fake.DeleteTeamStub = nil
	fake.deleteTeamReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) DeleteTeamReturnsOnCall(i int, result1 error) {
	fake.deleteTeamMutex.Lock()
	defer fake.deleteTeamMutex.Unlock()
	fake.DeleteTeamStub = nil
	if fake.deleteTeamReturnsOnCall == nil {
		fake.deleteTeamReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteTeamReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) Finalize(arg1 context.Context, arg2 db.Build) error {
	fake.finalizeMutex.Lock()
	ret, specificReturn := fake.finalizeReturnsOnCall[len(fake.finalizeArgsForCall)]
	fake.finalizeArgsForCall = append(fake.finalizeArgsForCall, struct {
		arg1 context.Context
		arg2 db.Build
	}{arg1, arg2})
	fake.recordInvocation("Finalize", []interface{}{arg1, arg2})
	fake.finalizeMutex.Unlock()
	if fake.FinalizeStub != nil {
		return fake.FinalizeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.finalizeReturns
	return fakeReturns.result1
}

func (fake *FakeEventStore) FinalizeCallCount() int {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	return len(fake.finalizeArgsForCall)
}

func (fake *FakeEventStore) FinalizeCalls(stub func(context.Context, db.Build) error) {
	fake.finalizeMutex.Lock()
	defer fake.finalizeMutex.Unlock()
	fake.FinalizeStub = stub
}

func (fake *FakeEventStore) FinalizeArgsForCall(i int) (context.Context, db.Build) {
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	argsForCall := fake.finalizeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventStore) FinalizeReturns(result1 error) {
	fake.finalizeMutex.Lock()
	defer fake.finalizeMutex.Unlock()
	fake.FinalizeStub = nil
	fake.finalizeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) FinalizeReturnsOnCall(i int, result1 error) {
	fake.finalizeMutex.Lock()
	defer fake.finalizeMutex.Unlock()
	fake.FinalizeStub = nil
	if fake.finalizeReturnsOnCall == nil {
		fake.finalizeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.finalizeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) Get(arg1 context.Context, arg2 db.Build, arg3 int, arg4 *db.EventKey) ([]event.Envelope, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 db.Build
		arg3 int
		arg4 *db.EventKey
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3, arg4})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEventStore) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeEventStore) GetCalls(stub func(context.Context, db.Build, int, *db.EventKey) ([]event.Envelope, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeEventStore) GetArgsForCall(i int) (context.Context, db.Build, int, *db.EventKey) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeEventStore) GetReturns(result1 []event.Envelope, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 []event.Envelope
		result2 error
	}{result1, result2}
}

func (fake *FakeEventStore) GetReturnsOnCall(i int, result1 []event.Envelope, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 []event.Envelope
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 []event.Envelope
		result2 error
	}{result1, result2}
}

func (fake *FakeEventStore) Initialize(arg1 context.Context, arg2 db.Build) error {
	fake.initializeMutex.Lock()
	ret, specificReturn := fake.initializeReturnsOnCall[len(fake.initializeArgsForCall)]
	fake.initializeArgsForCall = append(fake.initializeArgsForCall, struct {
		arg1 context.Context
		arg2 db.Build
	}{arg1, arg2})
	fake.recordInvocation("Initialize", []interface{}{arg1, arg2})
	fake.initializeMutex.Unlock()
	if fake.InitializeStub != nil {
		return fake.InitializeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.initializeReturns
	return fakeReturns.result1
}

func (fake *FakeEventStore) InitializeCallCount() int {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	return len(fake.initializeArgsForCall)
}

func (fake *FakeEventStore) InitializeCalls(stub func(context.Context, db.Build) error) {
	fake.initializeMutex.Lock()
	defer fake.initializeMutex.Unlock()
	fake.InitializeStub = stub
}

func (fake *FakeEventStore) InitializeArgsForCall(i int) (context.Context, db.Build) {
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	argsForCall := fake.initializeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventStore) InitializeReturns(result1 error) {
	fake.initializeMutex.Lock()
	defer fake.initializeMutex.Unlock()
	fake.InitializeStub = nil
	fake.initializeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) InitializeReturnsOnCall(i int, result1 error) {
	fake.initializeMutex.Lock()
	defer fake.initializeMutex.Unlock()
	fake.InitializeStub = nil
	if fake.initializeReturnsOnCall == nil {
		fake.initializeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initializeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) Put(arg1 context.Context, arg2 db.Build, arg3 []atc.Event) (db.EventKey, error) {
	var arg3Copy []atc.Event
	if arg3 != nil {
		arg3Copy = make([]atc.Event, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.putMutex.Lock()
	ret, specificReturn := fake.putReturnsOnCall[len(fake.putArgsForCall)]
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 context.Context
		arg2 db.Build
		arg3 []atc.Event
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("Put", []interface{}{arg1, arg2, arg3Copy})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.putReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEventStore) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeEventStore) PutCalls(stub func(context.Context, db.Build, []atc.Event) (db.EventKey, error)) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = stub
}

func (fake *FakeEventStore) PutArgsForCall(i int) (context.Context, db.Build, []atc.Event) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	argsForCall := fake.putArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeEventStore) PutReturns(result1 db.EventKey, result2 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 db.EventKey
		result2 error
	}{result1, result2}
}

func (fake *FakeEventStore) PutReturnsOnCall(i int, result1 db.EventKey, result2 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	if fake.putReturnsOnCall == nil {
		fake.putReturnsOnCall = make(map[int]struct {
			result1 db.EventKey
			result2 error
		})
	}
	fake.putReturnsOnCall[i] = struct {
		result1 db.EventKey
		result2 error
	}{result1, result2}
}

func (fake *FakeEventStore) Setup(arg1 context.Context) error {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Setup", []interface{}{arg1})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setupReturns
	return fakeReturns.result1
}

func (fake *FakeEventStore) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *FakeEventStore) SetupCalls(stub func(context.Context) error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = stub
}

func (fake *FakeEventStore) SetupArgsForCall(i int) context.Context {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	argsForCall := fake.setupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEventStore) SetupReturns(result1 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) SetupReturnsOnCall(i int, result1 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) UnmarshalKey(arg1 []byte, arg2 *db.EventKey) error {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.unmarshalKeyMutex.Lock()
	ret, specificReturn := fake.unmarshalKeyReturnsOnCall[len(fake.unmarshalKeyArgsForCall)]
	fake.unmarshalKeyArgsForCall = append(fake.unmarshalKeyArgsForCall, struct {
		arg1 []byte
		arg2 *db.EventKey
	}{arg1Copy, arg2})
	fake.recordInvocation("UnmarshalKey", []interface{}{arg1Copy, arg2})
	fake.unmarshalKeyMutex.Unlock()
	if fake.UnmarshalKeyStub != nil {
		return fake.UnmarshalKeyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.unmarshalKeyReturns
	return fakeReturns.result1
}

func (fake *FakeEventStore) UnmarshalKeyCallCount() int {
	fake.unmarshalKeyMutex.RLock()
	defer fake.unmarshalKeyMutex.RUnlock()
	return len(fake.unmarshalKeyArgsForCall)
}

func (fake *FakeEventStore) UnmarshalKeyCalls(stub func([]byte, *db.EventKey) error) {
	fake.unmarshalKeyMutex.Lock()
	defer fake.unmarshalKeyMutex.Unlock()
	fake.UnmarshalKeyStub = stub
}

func (fake *FakeEventStore) UnmarshalKeyArgsForCall(i int) ([]byte, *db.EventKey) {
	fake.unmarshalKeyMutex.RLock()
	defer fake.unmarshalKeyMutex.RUnlock()
	argsForCall := fake.unmarshalKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventStore) UnmarshalKeyReturns(result1 error) {
	fake.unmarshalKeyMutex.Lock()
	defer fake.unmarshalKeyMutex.Unlock()
	fake.UnmarshalKeyStub = nil
	fake.unmarshalKeyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) UnmarshalKeyReturnsOnCall(i int, result1 error) {
	fake.unmarshalKeyMutex.Lock()
	defer fake.unmarshalKeyMutex.Unlock()
	fake.UnmarshalKeyStub = nil
	if fake.unmarshalKeyReturnsOnCall == nil {
		fake.unmarshalKeyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unmarshalKeyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEventStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deletePipelineMutex.RLock()
	defer fake.deletePipelineMutex.RUnlock()
	fake.deleteTeamMutex.RLock()
	defer fake.deleteTeamMutex.RUnlock()
	fake.finalizeMutex.RLock()
	defer fake.finalizeMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.initializeMutex.RLock()
	defer fake.initializeMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.unmarshalKeyMutex.RLock()
	defer fake.unmarshalKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventStore) recordInvocation(key string, args []interface{}) {
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

var _ db.EventStore = new(FakeEventStore)