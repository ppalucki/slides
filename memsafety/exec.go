// +build OMIT
package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
	var err error

	cmd := exec.Command("sh", "-c", "while sleep 1; do date; done")
	// out := &bytes.Buffer{}
	// cmd.Stdout = out
	out, err := cmd.StdoutPipe() // it looks like a constuctor but actually, the owner is still owner of cmd object
	if err != nil {
		panic(err)
	}
	err = cmd.Start()

	// new pointer copying
	process := cmd.Process

	if err != nil {
		panic(err)
	}
	go func() {
		for {
			// time.Sleep(1 * time.Second)
			_, err := io.Copy(os.Stdout, out)
			if err != nil {
				panic(err)
			}
			// fmt.Println("n", n)
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		process.Kill()
	}()

	// fmt.Println(out.String())
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd.ProcessState)
}
