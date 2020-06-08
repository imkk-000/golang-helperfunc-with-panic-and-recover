package util_test

import (
	"errors"
	"testing"
	"util"
)

func TestFuncHelperRunWithSuccess(t *testing.T) {
	expectedStatusExecution := error(nil)
	helperFunc := util.FuncHelper{}
	helperFunc.New(func() {
		return
	})
	helperFunc.Run()

	actualStatusExecution := helperFunc.GetStatusExecution()

	if expectedStatusExecution != actualStatusExecution {
		t.Errorf("expect (%v) but it got (%v)", expectedStatusExecution, actualStatusExecution)
	}
}

func TestFuncHelperRunWithErrorPanic(t *testing.T) {
	expectedStatusExecution := errors.New("HelperFunc with panic")
	helperFunc := util.FuncHelper{}
	helperFunc.New(func() {
		panic(errors.New("HelperFunc with panic"))
	})
	helperFunc.Run()

	actualStatusExecution := helperFunc.GetStatusExecution()

	if expectedStatusExecution.Error() != actualStatusExecution.Error() {
		t.Errorf("expect (%v) but it got (%v)", expectedStatusExecution, actualStatusExecution)
	}
}

func TestFuncHelperRunWithErrorNoIndex(t *testing.T) {
	expectedStatusExecution := errors.New("runtime error: index out of range [2] with length 2")
	helperFunc := util.FuncHelper{}
	helperFunc.New(func() {
		v := []int{1, 2}
		for i := 0; i <= len(v); i++ {
			v[i]++
		}
	})
	helperFunc.Run()

	actualStatusExecution := helperFunc.GetStatusExecution()

	if expectedStatusExecution.Error() != actualStatusExecution.Error() {
		t.Errorf("expect (%v) but it got (%v)", expectedStatusExecution, actualStatusExecution)
	}
}
