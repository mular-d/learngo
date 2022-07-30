package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	var s = f()
	fmt.Println(s)

	v := 1
	increment(&v)
	fmt.Println(increment(&v))
}

func f() *int {
	v := 1
	return &v
}

func increment(p *int) int {
	*p++
	return *p
}
