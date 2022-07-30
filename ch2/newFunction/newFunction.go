package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	q := new(int)
	r := new(int)

	fmt.Println(q == r)
}
