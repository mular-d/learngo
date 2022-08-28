package main

import "fmt"

func main() {
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Tuesday)

	const (
		_ = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
		PiB
		EiB
		ZiB
		YiB
	)

	fmt.Println("1 MiB is", MiB, "Bytes")
}
