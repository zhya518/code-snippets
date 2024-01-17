package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("vim-go")
	capturingOutput2()
}

func base() {
	cmd := exec.Command("ls", "/usr/local/bin")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func sharingStandardOutput() {
	cmd := exec.Command("ls", "/usr/local/bin")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

// $ ls /usr/local/bin >output.log 2>&1
func capturingOutput() error {
	log, err := os.Create("output.log")
	if err != nil {
		return err
	}
	defer log.Close()
	cmd := exec.Command("ls", "/usr/local/bin")
	cmd.Stdout = log
	cmd.Stderr = log
	return cmd.Run()
}
func capturingOutput1() error {
	buf := new(bytes.Buffer)
	cmd := exec.Command("ls", "/usr/local/bin")
	cmd.Stdout = buf
	err := cmd.Run()
	if err != nil {
		return err
	}
	_ = buf
	//ProcessBuffer(buf)
	return nil
}

// One caveat of doing that is that if the spawned command generates a lot of output, the Buffer can grow without bound until the spawning process runs out of memory. You can process the output incrementally by using *Cmd.StdoutPipe(). You should call StdoutPipe before you start or run the process. StdoutPipe will return an io.ReadCloser which you can read from incrementally. It will be closed by os/exec when the spawned process terminates, at which point it is safe to call Wait().
func capturingOutput2() error {
	cmd := exec.Command("ls", "/usr/local/bin")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(stdout)
	err = cmd.Start()
	if err != nil {
		return err
	}
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		// Do something with the line here.
		//ProcessLine(scanner.Text())
	}
	if scanner.Err() != nil {
		cmd.Process.Kill()
		cmd.Wait()
		return scanner.Err()
	}
	return cmd.Wait()
}

// $ ls /usr/local/bin | grep pip
func pipingBetweenProcesses() error {
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

func errgroupAndCommandContext() error {
	eg, ctx := errgroup.WithContext(context.Background())
	sleeps := make([]*exec.Cmd, 3)
	sleeps[0] = exec.CommandContext(ctx, "sleep", "100")
	sleeps[1] = exec.CommandContext(ctx, "sleep", "100")
	sleeps[2] = exec.CommandContext(ctx, "sleep", "notanumber")
	for _, s := range sleeps {
		s := s
		eg.Go(func() error {
			return s.Run()
		})
	}
	return eg.Wait()
}

func processGroupsAndGracefulShutdown() error {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigch)

	proxy := exec.Command("bash", "-c", `
trap "echo proxy exiting" EXIT
echo "proxy started"
sleep 100
`)
	proxy.Stdout = os.Stdout
	proxy.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	server := exec.Command("bash", "-c", `
trap "echo server exiting" EXIT
echo server started
sleep 100
`)
	server.Stdout = os.Stdout
	server.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	err := proxy.Start()
	if err != nil {
		return err
	}
	defer proxy.Wait()

	err = server.Start()
	if err != nil {
		proxy.Process.Kill()
		return err
	}

	go func() {
		_, ok := <-sigch
		if ok {
			server.Process.Signal(syscall.SIGTERM)
			server.Wait()
			proxy.Process.Signal(syscall.SIGTERM)
		}
	}()

	return nil
}
