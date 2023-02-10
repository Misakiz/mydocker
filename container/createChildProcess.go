package container

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func CreateChildProcess(args []string) (error, string) {
	// args len equal one like [xx yy zz], so here need to split
	cmdArr := strings.Split(args[0], " ")
	//获取容器名字
	containerName := cmdArr[0]

	rootFolderPath := filepath.Join(ROOT_FOLDER_PATH_PREFEX, containerName, ROOTFS_NAME)

	//设置host
	if err := syscall.Sethostname([]byte(containerName)); err != nil {
		return err, "007"
	}
	//文件系统隔离
	if err := syscall.Chroot(rootFolderPath); err != nil {
		return err, "008"
	}
	//初始目录为/
	if err := syscall.Chdir("/"); err != nil {
		return err, "009"
	}
	//实现pid隔离 挂载proc
	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		return err, "010"
	}
	//找到命令绝对路径
	path, err := exec.LookPath(cmdArr[1])
	if err != nil {
		return err, "011"
	}

	//子进程替换父进程，为的避免是ps中出现父进程docker run xxx
	if err := syscall.Exec(path, cmdArr[1:], os.Environ()); err != nil {
		return err, "012"
	}
	return nil, ""
}
