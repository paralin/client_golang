// +build js wasm

package prometheus

// runtimeGetThreadCount returns the thread count.
func runtimeGetThreadCount() float64 {
	return 1 // 1 thread on js and wasm
}

// NewPidFileFn returns a function that retrieves a pid from the specified file.
// It is meant to be used for the PidFn field in ProcessCollectorOpts.
func NewPidFileFn(pidFilePath string) func() (int, error) {
	return func() (int, error) {
		// js stub: return empty
		return 0, nil
	}
}
