package main

import (
	"fmt"
	"math"
	"time"
)

// Do is a function x^e multiplied by a factor of 0.1 second.
// Result is limited to 2 minute.
func Do(attempts int) time.Duration {
	if attempts > 13 {
		return 2 * time.Minute
	}
	return time.Duration(math.Pow(float64(attempts), math.E)) * time.Millisecond * 100
}

func main() {
	fmt.Println("vim-go")
	for i := 0; i <= 100; i++ {
		fmt.Printf("i:%03d, result:%v\n", i, Do(i))
	}
}
