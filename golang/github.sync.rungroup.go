package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/uw-labs/sync/gogroup"
	"github.com/uw-labs/sync/rungroup"
)

// Extension of types from golang.org/x/sync.

func main() {
	fmt.Println("---runGroup---")
	runGroup()
	fmt.Println("---goGroup---")
	goGroup()
}

func goGroup() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	//Another alternative to error group that only waits until any single function started in the group terminates. Like rungroup it can only be created with a context and this context is cancelled as soon as any function started in the group terminates.
	//NOTE: calling wait without starting any goroutine with the Go method will block until the parent context is canceled.
	g, ctx := gogroup.New(ctx)
	g.Go(func() error {
		return run(ctx, time.Second)
	})
	g.Go(func() error {
		time.Sleep(time.Millisecond * 50)
		return errors.New("component stopped")
	})
	err := g.Wait()
	fmt.Println(err)
}

func runGroup() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Alternative to error group that stops (i.e. cancels the underlying context) as soon as any function started in the group terminates. For this to work it can only be created with a context.
	g, ctx := rungroup.New(ctx)

	g.Go(func() error {
		return run(ctx, time.Second)
	})
	g.Go(func() error {
		time.Sleep(time.Millisecond * 50)
		return errors.New("component stopped")
	})

	err := g.Wait()
	fmt.Println(err)
	//ctx done
	//component stopped
}

func run(ctx context.Context, d time.Duration) error {
	select {
	case <-ctx.Done():
		fmt.Println("ctx done")
		return ctx.Err()
	case <-time.After(d):
		fmt.Println("run")
		return nil
	}
}
