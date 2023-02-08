package run

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
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
	//创建SysProcAttr对象，为生产的进程设置uts隔离
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	//将os  /dev/stdin的标准输入输出 给cmd
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}

}
