// +build js wasm

package prometheus

// runtimeGetThreadCount returns the thread count.
func runtimeGetThreadCount() float64 {
	return 1 // 1 thread on js and wasm
}
