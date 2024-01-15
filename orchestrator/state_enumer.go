// Code generated by "enumer -type=State state.go"; DO NOT EDIT.

package orchestrator

import (
	"fmt"
	"strings"
)

const _StateName = "UnknownIdleRunningFatalError"

var _StateIndex = [...]uint8{0, 7, 11, 18, 28}

const _StateLowerName = "unknownidlerunningfatalerror"

func (i State) String() string {
	if i < 0 || i >= State(len(_StateIndex)-1) {
		return fmt.Sprintf("State(%d)", i)
	}
	return _StateName[_StateIndex[i]:_StateIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _StateNoOp() {
	var x [1]struct{}
	_ = x[Unknown-(0)]
	_ = x[Idle-(1)]
	_ = x[Running-(2)]
	_ = x[FatalError-(3)]
}

var _StateValues = []State{Unknown, Idle, Running, FatalError}

var _StateNameToValueMap = map[string]State{
	_StateName[0:7]:        Unknown,
	_StateLowerName[0:7]:   Unknown,
	_StateName[7:11]:       Idle,
	_StateLowerName[7:11]:  Idle,
	_StateName[11:18]:      Running,
	_StateLowerName[11:18]: Running,
	_StateName[18:28]:      FatalError,
	_StateLowerName[18:28]: FatalError,
}

var _StateNames = []string{
	_StateName[0:7],
	_StateName[7:11],
	_StateName[11:18],
	_StateName[18:28],
}

// StateString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func StateString(s string) (State, error) {
	if val, ok := _StateNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _StateNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to State values", s)
}

// StateValues returns all values of the enum
func StateValues() []State {
	return _StateValues
}

// StateStrings returns a slice of all String values of the enum
func StateStrings() []string {
	strs := make([]string, len(_StateNames))
	copy(strs, _StateNames)
	return strs
}

// IsAState returns "true" if the value is listed in the enum definition. "false" otherwise
func (i State) IsAState() bool {
	for _, v := range _StateValues {
		if i == v {
			return true
		}
	}
	return false
}
