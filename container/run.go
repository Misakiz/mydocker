package container

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

//
//func MyDockerRun() {
//	fmt.Println(os.Args)
//	fmt.Printf("Process -> %v [%d]\n", os.Args, os.Getpid())
//	if len(os.Args) <= 1 {
//		panic("!")
//	}
//	switch os.Args[1] {
//	case "container":
//		run()
//	case "init":
//		Init()
//	default:
//		panic("have no define")
//	}
//}
//
//func run() {
//	//开了一个子进程
//	cmd := exec.Command(os.Args[0], "init", os.Args[2])
//	//创建SysProcAttr对象，为生产的进程设置uts隔离
//	cmd.SysProcAttr = &syscall.SysProcAttr{
//		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC | syscall.CLONE_NEWUSER,
//		UidMappings: []syscall.SysProcIDMap{
//			{
//				0,
//				os.Getuid(),
//				1,
//			},
//		},
//		GidMappings: []syscall.SysProcIDMap{
//			{
//				0,
//				os.Getgid(),
//				1,
//			},
//		},
//	}
//
//	//将os  /dev/stdin的标准输入输出 给cmd
//	cmd.Stdin = os.Stdin
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//
//	if err := cmd.Start(); err != nil {
//		panic(err)
//	}
//
//	cmd.Wait()
//
//}
//
//func child() {
//	cmd := exec.Command(os.Args[2])
//	//设置host
//	syscall.Sethostname([]byte("container"))
//	// MS_NOEXEC: 在本文件系统中不允许运行其他程序
//	// MS_NOSUID: 在本系统中运行程序的时候，不允许set-user-id和不允许set-group-id
//	// MS_NODEV: 所有mount的系统都会默认设定的参数
//	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
//	//在子shell中mount proc目录，但未做文件系统隔离。
//	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
//	cmd.Stdin = os.Stdin
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	if err := cmd.Run(); err != nil {
//		panic(err)
//	}
//	syscall.Unmount("/proc", 0)
//}
//
//func Init() {
//	imageFolderPath := "/docker/images/base"
//	rootFolderPath := "/docker/containers/" + GenerateContainerId(64)
//
//	if _, err := os.Stat(rootFolderPath); os.IsNotExist(err) {
//		if err := CopyFileOrDirectory(imageFolderPath, rootFolderPath); err != nil {
//			panic(err)
//		}
//	}
//
//	//设置host
//	syscall.Sethostname([]byte("container"))
//
//	//设置文件系统隔离  相当于切到了一个空的目录中，空目录没有任何指令集
//	syscall.Chroot(rootFolderPath)
//	syscall.Chdir("/")
//	// MS_NOEXEC: 在本文件系统中不允许运行其他程序
//	// MS_NOSUID: 在本系统中运行程序的时候，不允许set-user-id和不允许set-group-id
//	// MS_NODEV: 所有mount的系统都会默认设定的参数
//	//defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
//	//在子shell中mount 父proc目录，但未做文件系统隔离。
//	syscall.Mount("proc", "/proc", "proc", 0, "")
//
//	//创建一个子进程并替换父进程环境
//	//Exec invokes the execve(2) system call.
//	//
//	//此方法会将在当前进程空间里，用新的程序覆盖掉当前程序，并执行新的程序，它们依然在同一个进程里，只是进程的内容发生了变化。
//	path, _ := exec.LookPath(os.Args[2])
//	if err := syscall.Exec(path, os.Args[2:], os.Environ()); err != nil {
//		panic(err)
//	}
//	syscall.Unmount("/proc", 0)
//}

func CopyFileOrDirectory(src string, dst string) error {
	//返回文件对象fileinfo
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	//判断src是否为目录
	if info.IsDir() {
		//如果为目录则在dst创建目录
		if err := os.MkdirAll(dst, 0777); err != nil {
			return err
		}
		//获取src目录的路径，并递归创建目录
		if list, err := ioutil.ReadDir(src); err == nil {
			for _, item := range list {
				//路径中的item   比如/xxx/bin 拷贝到    dst/bin
				if err = CopyFileOrDirectory(filepath.Join(src, item.Name()), filepath.Join(dst, item.Name())); err != nil {
					return err
				}
			}
			//
		} else {
			return err
		}
	} else {
		//如果不是路径，读取src文件流
		content, err := ioutil.ReadFile(src)
		if err != nil {
			return err
		}
		//函数向filename指定的文件中写入数据。
		if err := ioutil.WriteFile(dst, content, 0777); err != nil {
			return err
		}
	}
	return nil
}

//func GenerateContainerId(n uint) string {
//	rand.Seed(time.Now().UnixNano())
//	const letters = "dgkldankgdnakngjdsnlj"
//	b := make([]byte, n)
//	length := len(letters)
//	for i := range b {
//		b[i] = letters[rand.Intn(length)]
//	}
//	return string(b)
//}
//
//func CreateChildProcess(args []string) (error, string) {
//	// args len equal one like [xx yy zz], so here need to split
//	cmdArr := strings.Split(args[0], " ")
//	containerName := cmdArr[0]
//	rootFolderPath := filepath.Join(ROOT_FOLDER_PATH_PREFEX, containerName, ROOTFS_NAME)
//
//	if err := syscall.Sethostname([]byte(containerName)); err != nil {
//		return err, "007"
//	}
//	if err := syscall.Chroot(rootFolderPath); err != nil {
//		return err, "008"
//	}
//	if err := syscall.Chdir("/"); err != nil {
//		return err, "009"
//	}
//	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
//		return err, "010"
//	}
//	path, err := exec.LookPath(cmdArr[1])
//	if err != nil {
//		return err, "011"
//	}
//	if err := syscall.Exec(path, cmdArr[1:], os.Environ()); err != nil {
//		return err, "012"
//	}
//	return nil, ""
//}
