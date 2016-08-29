package common

import (
	"testing"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/packer"
	//"github.com/mitchellh/packer/packer"
)

type TestUi struct {
	Ui       packer.Ui
	uiCalled bool
	askCount int
}

func (u *TestUi) Ask(query string) (string, error) {
	u.askCount++
	return "", nil
}

func (u *TestUi) Say(message string) {}

func (u *TestUi) Message(message string) {}

func (u *TestUi) Error(message string) {}

func (u *TestUi) Machine(t string, args ...string) {}

func TestMultistepDebug_OnStep(t *testing.T) {
	ui := &TestUi{}
	pauseFn := common.MultistepDebugFn(ui)

	state := new(multistep.BasicStateBag)
	state.Put("debug", true)
	state.Put("debug_mode", packer.DebugOnStep)

	pauseFn(multistep.DebugLocationAfterRun, "SomeStep", state)
	pauseFn(multistep.DebugLocationAfterRun, "SomeStep", state)

	if ui.askCount != 2 {
		t.Fatalf("ui should be called twice")
	}
}

func TestMultistepDebug_OnError(t *testing.T) {
	ui := &TestUi{}
	pauseFn := common.MultistepDebugFn(ui)

	state := new(multistep.BasicStateBag)
	state.Put("debug", true)
	state.Put("debug_mode", packer.DebugOnError)

	state.Put("error", true)

	pauseFn(multistep.DebugLocationAfterRun, "SomeStep", state)
	pauseFn(multistep.DebugLocationAfterRun, "SomeStep", state)

	if ui.askCount != 1 {
		t.Fatalf("ui should be called once")
	}
}

func TestMultistepDebug_OnErrorWithoutError(t *testing.T) {
	ui := &TestUi{}
	pauseFn := common.MultistepDebugFn(ui)

	state := new(multistep.BasicStateBag)
	state.Put("debug", true)
	state.Put("debug_mode", packer.DebugOnError)

	state.Put("error", nil)

	pauseFn(multistep.DebugLocationAfterRun, "SomeStep", state)

	if ui.askCount != 0 {
		t.Fatalf("ui should not be called")
	}
}

func TestMultistepDebug_OnProvisioned(t *testing.T) {
	ui := &TestUi{}
	pauseFn := common.MultistepDebugFn(ui)

	state := new(multistep.BasicStateBag)
	state.Put("debug", true)
	state.Put("debug_mode", packer.DebugOnProvisioned)

	pauseFn(multistep.DebugLocationAfterRun, "StepProvision", state)
	pauseFn(multistep.DebugLocationAfterRun, "StepProvision", state)

	if ui.askCount != 1 {
		t.Fatalf("ui should be called once")
	}
}

func TestMultistepDebug_OnErrorOnProvisionedWithError(t *testing.T) {
	ui := &TestUi{}
	pauseFn := common.MultistepDebugFn(ui)

	state := new(multistep.BasicStateBag)
	state.Put("debug", true)
	state.Put("debug_mode", packer.DebugOnError|packer.DebugOnProvisioned)

	state.Put("error", true)

	pauseFn(multistep.DebugLocationAfterRun, "SomeStep", state)

	if ui.askCount != 1 {
		t.Fatalf("ui should be called once")
	}
}

func TestMultistepDebug_OnErrorOnProvisionedWithoutError(t *testing.T) {
	ui := &TestUi{}
	pauseFn := common.MultistepDebugFn(ui)

	state := new(multistep.BasicStateBag)
	state.Put("debug", true)
	state.Put("debug_mode", packer.DebugOnError|packer.DebugOnProvisioned)

	state.Put("error", nil)

	pauseFn(multistep.DebugLocationAfterRun, "SomeStep", state)
	pauseFn(multistep.DebugLocationAfterRun, "StepProvision", state)

	if ui.askCount != 1 {
		t.Fatalf("ui should be called once %d", ui.askCount)
	}
}
