package garden

import (
	"context"
	"runtime"
)

const NUM_BUTTERFLIES = 10

func watchButterflies(id int, x string) string {
	return x
}

func Garden(ctx context.Context) {
	butterflies(ctx)
	// Set log point in watchButterflies to see what they
	// are trying to tell you.
	runtime.Breakpoint()
}
