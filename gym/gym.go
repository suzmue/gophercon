package gym

import (
	"runtime"
)

func Gym() {
	searchLockers(lockers)
}

func searchLockers(lockers []string) {
	_ = lockers
	runtime.Breakpoint()
}
