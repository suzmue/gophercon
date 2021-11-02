package beach

import (
	"context"
	"fmt"
	"os"
	"runtime"
)

type hole struct {
	deeper *hole
	box    bool
}

func Beach(context context.Context) {
	for {
		runtime.Breakpoint()
		found := Pier(3)
		if found {
			fmt.Println("You won!")
			os.Exit(0)
		}
	}
}
func Pier(depth int) bool {
	hole := sandyspot
	for i := 0; i < depth; i++ {
		if hole == nil {
			break
		}
		hole = hole.deeper
	}
	if hole == nil {
		return false
	}
	return hole.box
}

func sidequest(password string) {
	if password == "gophercon" {
		fmt.Println("Here is a special clue!")
	}
}
