package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

//https://pkg.go.dev/golang.org/x/time/rate

func main() {
	// rate.Every 方法来指定向 Token 桶中放置 Token 的间隔
	limit := rate.Every(100 * time.Millisecond)
	_ = rate.NewLimiter(limit, 1)
	// 第一个参数是 r Limit。代表每秒可以向 Token 桶中产生多少 token。Limit 实际上是 float64 的别名
	// 第二个参数是 b int。b 代表 Token 桶的容量大小。
	// limiter1 := rate.NewLimiter(rate.Limit(2), 1)

	// func (lim *Limiter) Wait(ctx context.Context) (err error)
	// func (lim *Limiter) WaitN(ctx context.Context, n int) (err error)
	// func (lim *Limiter) Allow() bool
	// func (lim *Limiter) AllowN(now time.Time, n int) bool
	// func (lim *Limiter) Reserve() *Reservation
	// func (lim *Limiter) ReserveN(now time.Time, n int) *Reservation

	// 动态调整速率
	// SetLimit(Limit) 改变放入 Token 的速率
	// SetBurst(int) 改变 Token 桶大小
}

func limit() {
	fmt.Println("run limit")
	rl := rate.NewLimiter(100, 1)

	for i := 0; i < 10; i++ {
		start := time.Now()
		_ = rl.Wait(context.Background())
		fmt.Printf("#%d, took: %v\n", i, time.Since(start))
	}
}

func burst() {
	fmt.Println("run burst")
	rl := rate.NewLimiter(100, 2)

	for i := 0; i < 10; i++ {
		start := time.Now()
		_ = rl.WaitN(context.Background(), 2) // return err if burst > 2
		fmt.Printf("#%d, took: %v\n", i, time.Since(start))
	}
}

func reservation() {
	fmt.Println("run reservation")
	rl := rate.NewLimiter(10, 100)
	rl.WaitN(context.Background(), 100)

	re := rl.ReserveN(time.Now(), 50)
	if !re.OK() {
		return
	}

	fmt.Println("delay:", re.Delay())
}
