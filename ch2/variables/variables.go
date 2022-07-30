package main

import (
	"os"
)

func main() {
	i := 100
	var boiling float64 = 100

	var names []string
	var err error

	j, k := 0, 1
	j, k = j, k
}

func file() {
	in, err := os.Open("filename.txt")
	if err != nil {
		return err
	}
	in.Close()

	out, err := os.Create("filename.txt")
}
