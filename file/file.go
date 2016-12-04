package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("sh", "-c", "for i in {1..3};do echo $i;sleep 1;done")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	go io.Copy(os.Stderr, stdout) // reads cmd.Stdout.Fd
	go cmd.Process.Kill()
	exitCode := cmd.Wait() // closes cmd.Stdout.Fd descriptor
	fmt.Println(exitCode)
}
