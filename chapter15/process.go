package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Getuid(), os.Getgid(), os.Geteuid())

	fmt.Println(os.Getpid(), os.Getppid())

	fmt.Println(os.Getwd())
	fmt.Println(os.Chdir("e:\\"))
	fmt.Println(os.Getwd())

	process, _ := os.FindProcess(os.Getppid())
	fmt.Println(process)

	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	sprocess, err := os.StartProcess("C:\\Windows\\System32\\notepad.exe", []string{}, attr)
	fmt.Println(sprocess, err)

	// 10s中后kill进程
	go func() {
		time.Sleep(10 * time.Second)
		sprocess.Kill()
	}()
	state, err := sprocess.Wait() //等待进程结束
	if err == nil {
		fmt.Println(state.ExitCode())
		fmt.Println(state.Exited())
		fmt.Println(state.Pid())
		fmt.Println(state.Success())
		fmt.Println(state.Sys())
		fmt.Println(state.SysUsage())
		fmt.Println(state.SystemTime())
		fmt.Println(state.UserTime())

		sprocess.Release() //释放进程资源信息
	} else {
		fmt.Println(err)
	}

}
