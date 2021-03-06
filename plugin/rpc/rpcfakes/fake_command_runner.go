// This file was generated by counterfeiter
package rpcfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/commandregistry"
	"code.cloudfoundry.org/cli/plugin/rpc"
)

type FakeCommandRunner struct {
	CommandStub        func([]string, commandregistry.Dependency, bool) error
	commandMutex       sync.RWMutex
	commandArgsForCall []struct {
		arg1 []string
		arg2 commandregistry.Dependency
		arg3 bool
	}
	commandReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCommandRunner) Command(arg1 []string, arg2 commandregistry.Dependency, arg3 bool) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.commandMutex.Lock()
	fake.commandArgsForCall = append(fake.commandArgsForCall, struct {
		arg1 []string
		arg2 commandregistry.Dependency
		arg3 bool
	}{arg1Copy, arg2, arg3})
	fake.recordInvocation("Command", []interface{}{arg1Copy, arg2, arg3})
	fake.commandMutex.Unlock()
	if fake.CommandStub != nil {
		return fake.CommandStub(arg1, arg2, arg3)
	} else {
		return fake.commandReturns.result1
	}
}

func (fake *FakeCommandRunner) CommandCallCount() int {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	return len(fake.commandArgsForCall)
}

func (fake *FakeCommandRunner) CommandArgsForCall(i int) ([]string, commandregistry.Dependency, bool) {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	return fake.commandArgsForCall[i].arg1, fake.commandArgsForCall[i].arg2, fake.commandArgsForCall[i].arg3
}

func (fake *FakeCommandRunner) CommandReturns(result1 error) {
	fake.CommandStub = nil
	fake.commandReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCommandRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCommandRunner) recordInvocation(key string, args []interface{}) {
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

var _ rpc.CommandRunner = new(FakeCommandRunner)
