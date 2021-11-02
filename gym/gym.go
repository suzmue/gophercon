package gym

import (
	"context"
	"runtime"
)

func Gym(ctx context.Context) {
	searchLockers(lockers)
}

func searchLockers(lockers []string) {
	runtime.Breakpoint()
}
