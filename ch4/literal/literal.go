package main

import "fmt"

type P struct{ x, y int }

func main() {

	fmt.Println(scale(P{1, 2}, 5))
}

func scale(p P, factor int) P {
	return P{p.x * factor, p.y * factor}
}
