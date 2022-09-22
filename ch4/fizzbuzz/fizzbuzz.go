package main

import "fmt"

func main() {
	var output string
	for counter := 1; counter <= 100; counter++ {
		if counter%3 == 0 {
			output = "Fizz"
		}
		if counter%5 == 0 {
			output += "Buzz"
		}
		if counter%3 != 0 && counter%5 != 0 {
			fmt.Printf("%d", counter)
		}
		fmt.Printf("%s\n", output)
		output = ""
	}
}
