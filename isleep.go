package isleep

import (
	"context"
	"time"
)

// Sleep function works similarly to time.Sleep, but has an ability
// to be interrupted. It accepts a context which can be used for interruption and
// sleep duration period. Sleep returns true if it returned because of elapsed duration period.
// In the case of preliminary interruption it returns false.
func Sleep(ctx context.Context, d time.Duration) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	timer := time.NewTimer(d)

	select {
	case <-ctx.Done():
		timer.Stop()
		return false
	case <-timer.C:
		return true
	}
}
