package thelper

import (
	"reflect"
	"testing"
)

// ErrNameCase is an error case where there's an error function and a name that that error should have
type ErrNameCase struct {
	InputDesc   string
	ErrFunc     func() error
	WantErrName string
}

// CheckErrName is a function that operates on a list of error name cases and evaluates them
func CheckErrName(ecs []ErrNameCase, t *testing.T) {
	for _, val := range ecs {
		err := val.ErrFunc()
		if err == nil {
			if val.WantErrName != "nil" {
				t.Errorf("Wanted (%s) error name, but got back (nil) for input (%s)\nError: (%s)", val.WantErrName, val.InputDesc, "nil")
			}
		} else {
			errType := reflect.TypeOf(err)
			if errType.Name() != val.WantErrName {
				t.Errorf("Wanted (%s) error name, but got back (%s) for input (%s)\nError: (%s)", val.WantErrName, errType.Name(), val.InputDesc, err.Error())
			}
		}
	}
}
