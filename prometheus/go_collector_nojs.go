// +build !js,!wasm

package prometheus

import "runtime"

// runtimeGetThreadCount returns the thread count.
func runtimeGetThreadCount() float64 {
	n, _ := runtime.ThreadCreateProfile(nil)
	return float64(n)
}
