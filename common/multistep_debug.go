package common

import (
	"fmt"
	"log"
	"time"

	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

// MultistepDebugFn will return a proper multistep.DebugPauseFn to
// use for debugging if you're using multistep in your builder.
func MultistepDebugFn(ui packer.Ui) multistep.DebugPauseFn {
	return func(loc multistep.DebugLocation, name string, state multistep.StateBag) {
		var locationString string
		switch loc {
		case multistep.DebugLocationAfterRun:
			locationString = "after run of"
		case multistep.DebugLocationBeforeCleanup:
			locationString = "before cleanup of"
		default:
			locationString = "at"
		}

		message := fmt.Sprintf(
			"Pausing %s step '%s'. Press enter to continue.",
			locationString, name)

		result := make(chan string, 1)

		go func() {

			debugMode, ok := state.Get("debug_mode").(int)

			if ok {
				debug := (debugMode & packer.DebugOnStep) != 0

				if state.Get("handled-first-debug-step") == nil {
					debug = debug || (((debugMode & packer.DebugOnError) != 0) && state.Get("error") != nil)
					debug = debug || (((debugMode & packer.DebugOnProvisioned) != 0) && name == "StepProvision")
				}

				if debug {
					if state.Get("handled-first-debug-step") == nil {
						state.Put("handled-first-debug-step", true)
					}

					line, err := ui.Ask(message)
					if err != nil {
						log.Printf("Error asking for input: %s", err)
					}

					result <- line
				}
			}

			result <- ""
		}()

		for {
			select {
			case <-result:
				return
			case <-time.After(100 * time.Millisecond):
				if _, ok := state.GetOk(multistep.StateCancelled); ok {
					return
				}
			}
		}
	}
}
