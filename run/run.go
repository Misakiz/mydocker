package run

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func MyDockerRun() {
	fmt.Println(os.Args)
	fmt.Printf("Process -> %v [%d]\n", os.Args, os.Getpid())
	if len(os.Args) <= 1 {
		panic("!")
	}
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("have no define")
	}
}

func run() {
	//开了一个子进程
	cmd := exec.Command(os.Args[0], append([]string{"child"}, os.Args[2])...)
	//创建SysProcAttr对象，为生产的进程设置uts隔离
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	//将os  /dev/stdin的标准输入输出 给cmd
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}

}

func child() {
	cmd := exec.Command(os.Args[2])
	//设置host
	syscall.Sethostname([]byte("container"))
	// MS_NOEXEC: 在本文件系统中不允许运行其他程序
	// MS_NOSUID: 在本系统中运行程序的时候，不允许set-user-id和不允许set-group-id
	// MS_NODEV: 所有mount的系统都会默认设定的参数
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	syscall.Unmount("/proc", 0)
}
