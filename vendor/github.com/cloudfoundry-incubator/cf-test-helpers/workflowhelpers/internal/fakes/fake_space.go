// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type FakeSpace struct {
	CreateStub              func()
	createMutex             sync.RWMutex
	createArgsForCall       []struct{}
	DestroyStub             func()
	destroyMutex            sync.RWMutex
	destroyArgsForCall      []struct{}
	ShouldRemainStub        func() bool
	shouldRemainMutex       sync.RWMutex
	shouldRemainArgsForCall []struct{}
	shouldRemainReturns     struct {
		result1 bool
	}
	shouldRemainReturnsOnCall map[int]struct {
		result1 bool
	}
	OrganizationNameStub        func() string
	organizationNameMutex       sync.RWMutex
	organizationNameArgsForCall []struct{}
	organizationNameReturns     struct {
		result1 string
	}
	organizationNameReturnsOnCall map[int]struct {
		result1 string
	}
	SpaceNameStub        func() string
	spaceNameMutex       sync.RWMutex
	spaceNameArgsForCall []struct{}
	spaceNameReturns     struct {
		result1 string
	}
	spaceNameReturnsOnCall map[int]struct {
		result1 string
	}
	QuotaNameStub        func() string
	quotaNameMutex       sync.RWMutex
	quotaNameArgsForCall []struct{}
	quotaNameReturns     struct {
		result1 string
	}
	quotaNameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpace) Create() {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct{}{})
	fake.recordInvocation("Create", []interface{}{})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		fake.CreateStub()
	}
}

func (fake *FakeSpace) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSpace) Destroy() {
	fake.destroyMutex.Lock()
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct{}{})
	fake.recordInvocation("Destroy", []interface{}{})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		fake.DestroyStub()
	}
}

func (fake *FakeSpace) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeSpace) ShouldRemain() bool {
	fake.shouldRemainMutex.Lock()
	ret, specificReturn := fake.shouldRemainReturnsOnCall[len(fake.shouldRemainArgsForCall)]
	fake.shouldRemainArgsForCall = append(fake.shouldRemainArgsForCall, struct{}{})
	fake.recordInvocation("ShouldRemain", []interface{}{})
	fake.shouldRemainMutex.Unlock()
	if fake.ShouldRemainStub != nil {
		return fake.ShouldRemainStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.shouldRemainReturns.result1
}

func (fake *FakeSpace) ShouldRemainCallCount() int {
	fake.shouldRemainMutex.RLock()
	defer fake.shouldRemainMutex.RUnlock()
	return len(fake.shouldRemainArgsForCall)
}

func (fake *FakeSpace) ShouldRemainReturns(result1 bool) {
	fake.ShouldRemainStub = nil
	fake.shouldRemainReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSpace) ShouldRemainReturnsOnCall(i int, result1 bool) {
	fake.ShouldRemainStub = nil
	if fake.shouldRemainReturnsOnCall == nil {
		fake.shouldRemainReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.shouldRemainReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSpace) OrganizationName() string {
	fake.organizationNameMutex.Lock()
	ret, specificReturn := fake.organizationNameReturnsOnCall[len(fake.organizationNameArgsForCall)]
	fake.organizationNameArgsForCall = append(fake.organizationNameArgsForCall, struct{}{})
	fake.recordInvocation("OrganizationName", []interface{}{})
	fake.organizationNameMutex.Unlock()
	if fake.OrganizationNameStub != nil {
		return fake.OrganizationNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.organizationNameReturns.result1
}

func (fake *FakeSpace) OrganizationNameCallCount() int {
	fake.organizationNameMutex.RLock()
	defer fake.organizationNameMutex.RUnlock()
	return len(fake.organizationNameArgsForCall)
}

func (fake *FakeSpace) OrganizationNameReturns(result1 string) {
	fake.OrganizationNameStub = nil
	fake.organizationNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpace) OrganizationNameReturnsOnCall(i int, result1 string) {
	fake.OrganizationNameStub = nil
	if fake.organizationNameReturnsOnCall == nil {
		fake.organizationNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.organizationNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpace) SpaceName() string {
	fake.spaceNameMutex.Lock()
	ret, specificReturn := fake.spaceNameReturnsOnCall[len(fake.spaceNameArgsForCall)]
	fake.spaceNameArgsForCall = append(fake.spaceNameArgsForCall, struct{}{})
	fake.recordInvocation("SpaceName", []interface{}{})
	fake.spaceNameMutex.Unlock()
	if fake.SpaceNameStub != nil {
		return fake.SpaceNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.spaceNameReturns.result1
}

func (fake *FakeSpace) SpaceNameCallCount() int {
	fake.spaceNameMutex.RLock()
	defer fake.spaceNameMutex.RUnlock()
	return len(fake.spaceNameArgsForCall)
}

func (fake *FakeSpace) SpaceNameReturns(result1 string) {
	fake.SpaceNameStub = nil
	fake.spaceNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpace) SpaceNameReturnsOnCall(i int, result1 string) {
	fake.SpaceNameStub = nil
	if fake.spaceNameReturnsOnCall == nil {
		fake.spaceNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.spaceNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpace) QuotaName() string {
	fake.quotaNameMutex.Lock()
	ret, specificReturn := fake.quotaNameReturnsOnCall[len(fake.quotaNameArgsForCall)]
	fake.quotaNameArgsForCall = append(fake.quotaNameArgsForCall, struct{}{})
	fake.recordInvocation("QuotaName", []interface{}{})
	fake.quotaNameMutex.Unlock()
	if fake.QuotaNameStub != nil {
		return fake.QuotaNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.quotaNameReturns.result1
}

func (fake *FakeSpace) QuotaNameCallCount() int {
	fake.quotaNameMutex.RLock()
	defer fake.quotaNameMutex.RUnlock()
	return len(fake.quotaNameArgsForCall)
}

func (fake *FakeSpace) QuotaNameReturns(result1 string) {
	fake.QuotaNameStub = nil
	fake.quotaNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpace) QuotaNameReturnsOnCall(i int, result1 string) {
	fake.QuotaNameStub = nil
	if fake.quotaNameReturnsOnCall == nil {
		fake.quotaNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.quotaNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpace) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.shouldRemainMutex.RLock()
	defer fake.shouldRemainMutex.RUnlock()
	fake.organizationNameMutex.RLock()
	defer fake.organizationNameMutex.RUnlock()
	fake.spaceNameMutex.RLock()
	defer fake.spaceNameMutex.RUnlock()
	fake.quotaNameMutex.RLock()
	defer fake.quotaNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpace) recordInvocation(key string, args []interface{}) {
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