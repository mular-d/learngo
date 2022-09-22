package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	delete(ages, "charlie")
	ages["bob"] = 35

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	if _, ok := ages["charlie"]; !ok {
		fmt.Println("Charlie is not found.")
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
