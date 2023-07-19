# Interruptible Sleep

This simple module implements a single `Sleep` function which is similar
to `time.Sleep()` from Go standard library, but can be interrupted. In order
to implement interruption standard context mechanism is used.

Example usage:
```go
ctx, cancel := context.WithCancel(parent)
res := isleep.Sleep(ctx, d)
if res {
	// Sleep finished due to elapsed period
} else {
	// Sleep finished due to interruption
}
```