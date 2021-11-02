package main

import "fmt"

func main() {
	fmt.Println("monkey running")
	var caught bool
	var x int
	for !caught {
		x += 1
	}
	catchTheMonkey()
}

func catchTheMonkey() {
	panic("The monkey panics as you corner him! He tells you that you need to see the bigger picture. (Reveal the global variables to see your next clue.)")
}
