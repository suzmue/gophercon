package jungle

import (
	"os"
	"os/exec"
)

func Guardian() {
	exec.Command("go", "build", "-o=/Users/suzmue/gophercon/monkey/monkey", "-gcflags=all=-N -l", "github.com/suzmue/gophercon21/monkey").Run()

	cmd := exec.Command("/Users/suzmue/gophercon/monkey/monkey")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
