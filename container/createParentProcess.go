package container

import (
	"mydocker/alert"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func CreateParentProcess(containerName string, interactive bool, tty bool, args []string) *exec.Cmd {
	//
	args = append([]string{containerName}, args[0:]...)
	//日志输入路径
	logFilePath := filepath.Join(ROOT_FOLDER_PATH_PREFEX, containerName, LOG_FILE_NAME)
	//创建SysProcAttr对象，为进程设置uts隔离

	//相当于执行docker child
	cmd := exec.Command("/proc/self/exe", "child", strings.Join(args, " "))
	//封装sysProc对象设置好进程隔离对象
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}
	imageFolderPath := IMAGE_FOLDER_PATH
	rootFolderPath := filepath.Join(ROOT_FOLDER_PATH_PREFEX, containerName, ROOTFS_NAME)
	//目录不存在则创建
	if _, err := os.Stat(rootFolderPath); os.IsNotExist(err) {
		if err := CopyFileOrDirectory(imageFolderPath, rootFolderPath); err != nil {
			alert.Show(err, "013")
		}
	}
	//是否有交互
	if tty {
		if interactive {
			cmd.Stdin = os.Stdin
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		// detach mode
		//创建日志目录
		logFile, err := os.Create(logFilePath)
		if err != nil {
			alert.Show(err, "014")
		}
		//将日志文件给stdout
		cmd.Stdout = logFile
		alert.Println(containerName)
	}
	return cmd
}
