package beach

import (
	"context"
	"runtime"
)

func Beach(_ context.Context) {
	runtime.Breakpoint()
	searchBeach()
}
