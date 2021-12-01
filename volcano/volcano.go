package volcano

import (
	"context"
	"runtime"
)

func Volcano(ctx context.Context) {
	// This is a dead end.
	runtime.Breakpoint()
}
