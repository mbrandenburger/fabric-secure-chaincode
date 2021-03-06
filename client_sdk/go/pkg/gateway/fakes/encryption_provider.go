// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/hyperledger/fabric-private-chaincode/internal/crypto"
)

type EncryptionProvider struct {
	NewEncryptionContextStub        func() (crypto.EncryptionContext, error)
	newEncryptionContextMutex       sync.RWMutex
	newEncryptionContextArgsForCall []struct {
	}
	newEncryptionContextReturns struct {
		result1 crypto.EncryptionContext
		result2 error
	}
	newEncryptionContextReturnsOnCall map[int]struct {
		result1 crypto.EncryptionContext
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EncryptionProvider) NewEncryptionContext() (crypto.EncryptionContext, error) {
	fake.newEncryptionContextMutex.Lock()
	ret, specificReturn := fake.newEncryptionContextReturnsOnCall[len(fake.newEncryptionContextArgsForCall)]
	fake.newEncryptionContextArgsForCall = append(fake.newEncryptionContextArgsForCall, struct {
	}{})
	stub := fake.NewEncryptionContextStub
	fakeReturns := fake.newEncryptionContextReturns
	fake.recordInvocation("NewEncryptionContext", []interface{}{})
	fake.newEncryptionContextMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EncryptionProvider) NewEncryptionContextCallCount() int {
	fake.newEncryptionContextMutex.RLock()
	defer fake.newEncryptionContextMutex.RUnlock()
	return len(fake.newEncryptionContextArgsForCall)
}

func (fake *EncryptionProvider) NewEncryptionContextCalls(stub func() (crypto.EncryptionContext, error)) {
	fake.newEncryptionContextMutex.Lock()
	defer fake.newEncryptionContextMutex.Unlock()
	fake.NewEncryptionContextStub = stub
}

func (fake *EncryptionProvider) NewEncryptionContextReturns(result1 crypto.EncryptionContext, result2 error) {
	fake.newEncryptionContextMutex.Lock()
	defer fake.newEncryptionContextMutex.Unlock()
	fake.NewEncryptionContextStub = nil
	fake.newEncryptionContextReturns = struct {
		result1 crypto.EncryptionContext
		result2 error
	}{result1, result2}
}

func (fake *EncryptionProvider) NewEncryptionContextReturnsOnCall(i int, result1 crypto.EncryptionContext, result2 error) {
	fake.newEncryptionContextMutex.Lock()
	defer fake.newEncryptionContextMutex.Unlock()
	fake.NewEncryptionContextStub = nil
	if fake.newEncryptionContextReturnsOnCall == nil {
		fake.newEncryptionContextReturnsOnCall = make(map[int]struct {
			result1 crypto.EncryptionContext
			result2 error
		})
	}
	fake.newEncryptionContextReturnsOnCall[i] = struct {
		result1 crypto.EncryptionContext
		result2 error
	}{result1, result2}
}

func (fake *EncryptionProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newEncryptionContextMutex.RLock()
	defer fake.newEncryptionContextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EncryptionProvider) recordInvocation(key string, args []interface{}) {
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
