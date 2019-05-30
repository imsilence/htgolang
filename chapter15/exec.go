package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {

	fmt.Println(exec.LookPath("notepad"))

	// 执行命令获取结果
	cmd := exec.Command("ping", "-n", "2", "www.baidu.com")
	fmt.Printf("%#v\n", cmd)
	output, _ := cmd.Output()
	fmt.Println(string(output))

	//执行命令
	cmd = exec.Command("ping", "-n", "2", "www.baidu.com")
	fmt.Println(cmd.Path, cmd.Args)
	fmt.Println(cmd.Run())

	//执行命令通过管道获取结果
	cmd = exec.Command("ping", "-n", "2", "www.baidu.com")
	stdout, _ := cmd.StdoutPipe()
	fmt.Println(cmd.Path, cmd.Args)

	fmt.Println(cmd.Start())

	io.Copy(os.Stdout, stdout)
	fmt.Println(cmd.Wait())

	// 执行命令，使用管道获取结果并输入给Json Decoder
	user := make(map[string]string)
	cmd = exec.Command("C:\\Program Files\\Git\\usr\\bin\\echo", `{"name" : "kk"}`)
	stdout, _ = cmd.StdoutPipe()

	cmd.Start()

	json.NewDecoder(stdout).Decode(&user)
	cmd.Wait()
	fmt.Printf("%#v\n", user)

	// 执行命令，使用管道连接多个执行的命令
	cmd01 := exec.Command("C:\\Program Files\\Git\\usr\\bin\\echo", `{"name" : "kk"}`)
	cmd02 := exec.Command("python", "-m", "json.tool")

	stdout01, _ := cmd01.StdoutPipe()
	cmd02.Stdin = stdout01

	stdout02, _ := cmd02.StdoutPipe()

	fmt.Println(cmd01.Start())
	fmt.Println(cmd02.Start())

	fmt.Println(cmd01.Wait())

	io.Copy(os.Stdout, stdout02)
	fmt.Println(cmd02.Wait())

}
