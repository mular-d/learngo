package main

import "fmt"

func main() {
	// var a [5]int = [5]int{1, 2, 3, 4, 5}
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println("First element ->", a[0])
	fmt.Println("Last element ->", a[len(a)-1])

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d -> %d\n", i, v)
	}

	// Print the elements only.
	fmt.Println("List of elements")
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "e", GBP: "f", RMB: "y"}
	fmt.Println(RMB, symbol[RMB])

	// Print 99 0's and -1 lastly
	r := [...]int{99: -1}
	for _, v := range r {
		fmt.Printf("%d\n", v)
	}
}
