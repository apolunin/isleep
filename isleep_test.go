package isleep

import (
	"context"
	"testing"
	"time"
)

func TestSleep_Regular(t *testing.T) {
	const period = 1 * time.Second

	if res := Sleep(context.Background(), period); !res {
		t.Errorf("expected result: %t, got: %t", true, res)
	}
}

func TestSleep_Interrupted(t *testing.T) {
	const period = 1 * time.Hour

	var (
		ctx, cancel = context.WithCancel(context.Background())
		start       = make(chan struct{})
	)

	go func() {
		<-start
		cancel()
	}()

	close(start)

	if res := Sleep(ctx, period); res {
		t.Errorf("expected result: %t, got: %t", false, res)
	}
}

func TestSleep_NilContext(t *testing.T) {
	const period = 1 * time.Second

	if res := Sleep(nil, period); !res {
		t.Errorf("expected result: %t, got: %t", true, res)
	}
}
