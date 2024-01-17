package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	pr, pw := io.Pipe()

	cmd := exec.Command("/bin/bash", "-c", "cat")
	cmd.Stdin = pr

	started := make(chan error)
	done := make(chan error)
	go func() {
		started <- cmd.Start()
		done <- cmd.Wait()
	}()

	<-started
	println("Started")
	pw.Write([]byte("hello"))

	cmd.Process.Signal(syscall.SIGTERM)

	err := <-done
	println("Done")
	if err != nil {
		fmt.Printf("wait err: %s\n", err.Error())
	}
}

func foo() {
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}
	defer r.Close()
	ls := exec.Command("ls", "/usr/local/bin")
	ls.Stdout = w
	err = ls.Start()
	if err != nil {
		return err
	}
	defer ls.Wait()
	w.Close()
	grep := exec.Command("grep", "pip")
	grep.Stdin = r
	grep.Stdout = os.Stdout
	return grep.Run()
}
