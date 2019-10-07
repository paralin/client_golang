// +build !js,!wasm

package prometheus

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
)

// runtimeGetThreadCount returns the thread count.
func runtimeGetThreadCount() float64 {
	n, _ := runtime.ThreadCreateProfile(nil)
	return float64(n)
}

// NewPidFileFn returns a function that retrieves a pid from the specified file.
// It is meant to be used for the PidFn field in ProcessCollectorOpts.
func NewPidFileFn(pidFilePath string) func() (int, error) {
	return func() (int, error) {
		// js stub: return empty
		content, err := ioutil.ReadFile(pidFilePath)
		if err != nil {
			return 0, fmt.Errorf("can't read pid file %q: %+v", pidFilePath, err)
		}
		pid, err := strconv.Atoi(strings.TrimSpace(string(content)))
		if err != nil {
			return 0, fmt.Errorf("can't parse pid file %q: %+v", pidFilePath, err)
		}

		return pid, nil
	}
}
