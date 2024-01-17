package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	hunch "github.com/aaronjan/hunch"
)

func main() {
	ctx := context.Background()
	r, err := hunch.All(
		ctx,
		func(ctx context.Context) (any, error) {
			fmt.Println("#1 start")
			time.Sleep(300 * time.Millisecond)
			return 1, nil
		},
		func(ctx context.Context) (any, error) {
			fmt.Println("#2 start")
			time.Sleep(200 * time.Millisecond)
			//return 2, nil
			return 0, errors.New("#2 failure")
		},
		func(ctx context.Context) (any, error) {
			fmt.Println("#3 start")
			time.Sleep(100 * time.Millisecond)
			return 3, nil
			//return 0, errors.New("#3 failure")
		},
	)

	fmt.Println(r, err)
}
