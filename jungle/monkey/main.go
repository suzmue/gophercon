package main

import "time"

func main() {
	bait := "banana"
	// Break out of the loop after a timeout.
	go func() {
		time.Sleep(5 * time.Minute)
		bait = i_want
	}()
	// showGlobalVariables to figure out
	// what to put in the trap.
	for bait != i_want {
		_ = bait
	}
	cornerTheMonkey()
}

func cornerTheMonkey() {
	// The monkey panics as you catch him in the trap.
	panic("The butterflies in the garden can tell you where to find the treasure.")
}
