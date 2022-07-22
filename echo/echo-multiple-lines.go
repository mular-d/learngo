package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
	fmt.Println(time.Since(start).Seconds())
}
