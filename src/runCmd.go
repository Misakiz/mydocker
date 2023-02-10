package main

import (
	"github.com/spf13/cobra"
	"math/rand"
	"mydocker/alert"
	"mydocker/container"
	"os"
	"time"
)

func InitRunCmd() *cobra.Command {
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run a command in a new container",
		Run: func(self *cobra.Command, args []string) {
			//是否以交互模式启动
			is_tty, err := self.Flags().GetBool("tty")
			if err != nil {
				alert.Show(err, "001")
			}
			is_interactive, err := self.Flags().GetBool("interactive")
			if err != nil {
				alert.Show(err, "002")
			}
			is_detach, err := self.Flags().GetBool("detach")
			if err != nil {
				alert.Show(err, "003")
			}
			//两个只能取其中一个
			if is_detach && is_tty {
				alert.Show(nil, "004")
				return
			}
			//判断是否有指定名字，有则获取容器名字
			containerName, err := self.Flags().GetString("name")
			if err != nil {
				alert.Show(err, "005")
			}
			//容器名字为空则随机生成一个值
			if containerName == "" {
				containerName = GenerateContainerId(container.MAX_CONTAINER_ID)
			}
			//创建父进程
			cmd := container.CreateParentProcess(containerName, is_interactive, is_tty, args)
			if err := cmd.Start(); err != nil {
				alert.Show(err, "006")
			}
			//父进程消失，其fork的子进程会被一号进程接管，实现-d执行原理
			if !is_detach {
				cmd.Wait()
			}
			os.Exit(-1)
		},
	}
	runCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	runCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	runCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")
	runCmd.Flags().StringP("name", "n", "", "Assign a name to the container")
	return runCmd
}

func GenerateContainerId(n uint) string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	length := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(length)]
	}
	return string(b)
}
