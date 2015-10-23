// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/api/containerserver"
	"github.com/concourse/atc/db"
)

type FakeContainerDB struct {
	GetContainerStub        func(handle string) (db.Container, bool, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		handle string
	}
	getContainerReturns struct {
		result1 db.Container
		result2 bool
		result3 error
	}
	FindContainersByIdentifierStub        func(db.ContainerIdentifier) ([]db.Container, error)
	findContainersByIdentifierMutex       sync.RWMutex
	findContainersByIdentifierArgsForCall []struct {
		arg1 db.ContainerIdentifier
	}
	findContainersByIdentifierReturns struct {
		result1 []db.Container
		result2 error
	}
}

func (fake *FakeContainerDB) GetContainer(handle string) (db.Container, bool, error) {
	fake.getContainerMutex.Lock()
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		handle string
	}{handle})
	fake.getContainerMutex.Unlock()
	if fake.GetContainerStub != nil {
		return fake.GetContainerStub(handle)
	} else {
		return fake.getContainerReturns.result1, fake.getContainerReturns.result2, fake.getContainerReturns.result3
	}
}

func (fake *FakeContainerDB) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeContainerDB) GetContainerArgsForCall(i int) string {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return fake.getContainerArgsForCall[i].handle
}

func (fake *FakeContainerDB) GetContainerReturns(result1 db.Container, result2 bool, result3 error) {
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 db.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeContainerDB) FindContainersByIdentifier(arg1 db.ContainerIdentifier) ([]db.Container, error) {
	fake.findContainersByIdentifierMutex.Lock()
	fake.findContainersByIdentifierArgsForCall = append(fake.findContainersByIdentifierArgsForCall, struct {
		arg1 db.ContainerIdentifier
	}{arg1})
	fake.findContainersByIdentifierMutex.Unlock()
	if fake.FindContainersByIdentifierStub != nil {
		return fake.FindContainersByIdentifierStub(arg1)
	} else {
		return fake.findContainersByIdentifierReturns.result1, fake.findContainersByIdentifierReturns.result2
	}
}

func (fake *FakeContainerDB) FindContainersByIdentifierCallCount() int {
	fake.findContainersByIdentifierMutex.RLock()
	defer fake.findContainersByIdentifierMutex.RUnlock()
	return len(fake.findContainersByIdentifierArgsForCall)
}

func (fake *FakeContainerDB) FindContainersByIdentifierArgsForCall(i int) db.ContainerIdentifier {
	fake.findContainersByIdentifierMutex.RLock()
	defer fake.findContainersByIdentifierMutex.RUnlock()
	return fake.findContainersByIdentifierArgsForCall[i].arg1
}

func (fake *FakeContainerDB) FindContainersByIdentifierReturns(result1 []db.Container, result2 error) {
	fake.FindContainersByIdentifierStub = nil
	fake.findContainersByIdentifierReturns = struct {
		result1 []db.Container
		result2 error
	}{result1, result2}
}

var _ containerserver.ContainerDB = new(FakeContainerDB)
