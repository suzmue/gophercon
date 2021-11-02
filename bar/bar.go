package bar

import (
	"context"
	"runtime"
)

func Bar(ctx context.Context) {
	enterBar()
}

func GatherIntel() {
	runtime.Breakpoint()
}
