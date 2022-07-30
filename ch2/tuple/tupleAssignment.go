package main

import "fmt"

func main() {
	x, y := 2, 3
	x, y = y, x

	fmt.Println(fib(20))
	fmt.Println(gcd(15, 20))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
