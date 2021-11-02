package sevenSeas

import "fmt"

func Sail() {

	crossChannel()
}

func crossChannel() {
	numberOfSpins, sailsDown := 0, true
	for whirlpool && sailsDown {
		numberOfSpins++
		if numberOfSpins%1000000 == 0 {
			fmt.Println("Help! I'm spinning out of control!")
		}
	}
	fmt.Println("Phew! Made it safely past the whirlpool :)")
}
