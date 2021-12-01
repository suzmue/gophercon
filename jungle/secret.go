package jungle

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func Guardian() {
	dir, err := os.MkdirTemp("", "treasureHunt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(dir)
	output := filepath.Join(dir, "monkey")
	if runtime.GOOS == "windows" {
		output += ".exe"
	}
	err = exec.Command("go", "build", fmt.Sprintf("-o=%s", output), "-gcflags=all=-N -l", "gophercon/jungle/monkey").Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(output)
	exec.Command(output).Run()
}
