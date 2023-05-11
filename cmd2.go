package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

func main() {
	// 要执行的命令
	//cmd := exec.Command("ls", "-lah")

	//cmd := exec.Command("bash.exe", "-c", "echo 111")

	stdout, stderr, code := RunCommand("ls", "-lah")
	fmt.Println(stdout, stderr, code)

}

func RunCommand(name string, args ...string) (stdout string, stderr string, exitCode int) {

	var defaultFailedCode = -9999

	log.Println("Run command:", name, args)

	var outbuf, errbuf bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout = outbuf.String()
	stderr = errbuf.String()

	if err != nil {
		//尝试获取exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			// 如果在这种情况下`name` 在$PATH 中不可用，就会发生这种情况（在 OSX 中），无法获取退出代码，
			//并且stderr 很可能是空字符串，因此我们使用默认的失败代码，并将err 格式化为字符串和设置为标准错误
			log.Printf("Could not get exit code for failed program: %v, %v", name, args)
			exitCode = defaultFailedCode
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		// 成功，exitCode 应该是 0
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	log.Printf("command result, stdout: %v, stderr: %v, exitCode: %v", stdout, stderr, exitCode)
	return
}
