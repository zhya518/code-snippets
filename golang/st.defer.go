package main

import "fmt"

func deferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func deferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func deferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func deferFunc4() (t int) {
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}

func main() {
	fmt.Println(deferFunc1(1)) // 4
	fmt.Println(deferFunc2(1)) // 1
	fmt.Println(deferFunc3(1)) // 3
	deferFunc4()               // 0 2
}
