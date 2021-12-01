package capitol

import (
	"context"
	"runtime"
)

func Capitol(ctx context.Context) {
	// This is a dead end.
	runtime.Breakpoint()
}
