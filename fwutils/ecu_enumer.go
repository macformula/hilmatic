// Code generated by "enumer -type=Ecu ecu.go"; DO NOT EDIT.

package fwutils

import (
	"fmt"
	"strings"
)

const _EcuName = "UnknownEcuFrontControllerLvControllerTmsControllerDashController"

var _EcuIndex = [...]uint8{0, 10, 25, 37, 50, 64}

const _EcuLowerName = "unknownecufrontcontrollerlvcontrollertmscontrollerdashcontroller"

func (i Ecu) String() string {
	if i < 0 || i >= Ecu(len(_EcuIndex)-1) {
		return fmt.Sprintf("Ecu(%d)", i)
	}
	return _EcuName[_EcuIndex[i]:_EcuIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _EcuNoOp() {
	var x [1]struct{}
	_ = x[UnknownEcu-(0)]
	_ = x[FrontController-(1)]
	_ = x[LvController-(2)]
	_ = x[TmsController-(3)]
	_ = x[DashController-(4)]
}

var _EcuValues = []Ecu{UnknownEcu, FrontController, LvController, TmsController, DashController}

var _EcuNameToValueMap = map[string]Ecu{
	_EcuName[0:10]:       UnknownEcu,
	_EcuLowerName[0:10]:  UnknownEcu,
	_EcuName[10:25]:      FrontController,
	_EcuLowerName[10:25]: FrontController,
	_EcuName[25:37]:      LvController,
	_EcuLowerName[25:37]: LvController,
	_EcuName[37:50]:      TmsController,
	_EcuLowerName[37:50]: TmsController,
	_EcuName[50:64]:      DashController,
	_EcuLowerName[50:64]: DashController,
}

var _EcuNames = []string{
	_EcuName[0:10],
	_EcuName[10:25],
	_EcuName[25:37],
	_EcuName[37:50],
	_EcuName[50:64],
}

// EcuString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func EcuString(s string) (Ecu, error) {
	if val, ok := _EcuNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _EcuNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Ecu values", s)
}

// EcuValues returns all values of the enum
func EcuValues() []Ecu {
	return _EcuValues
}

// EcuStrings returns a slice of all String values of the enum
func EcuStrings() []string {
	strs := make([]string, len(_EcuNames))
	copy(strs, _EcuNames)
	return strs
}

// IsAEcu returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Ecu) IsAEcu() bool {
	for _, v := range _EcuValues {
		if i == v {
			return true
		}
	}
	return false
}
