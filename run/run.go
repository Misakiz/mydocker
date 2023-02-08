package ran

import (
	"fmt"
	"os"
	"os/exec"
)

func MyDockerRun() {
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))
	if len(os.Args) <= 1 {
		panic("!")
	}
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("have no define")
	}
}

func run() {
	cmd := exec.Command(os.Args[2])
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
