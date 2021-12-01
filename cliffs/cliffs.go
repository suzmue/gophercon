package cliffs

import (
	"context"
	"runtime"
)

func Cliffs(ctx context.Context) {
	// This is a dead end.
	runtime.Breakpoint()
}
