package main

import "fmt"

func main() {
	// s := "hello, world"
	// fmt.Println(len(s))
	// fmt.Println(s[0], s[7])

	// c := s[len(s)] // panic: index out of range
	// fmt.Println(c)
	// fmt.Println(s)
	// fmt.Println(s[0:5])
	// fmt.Println("goodbye" + s[5:])

	// s[0] = 'L'
	fmt.Println(HasPrefix("Hello world", "hello"))
	fmt.Println(HasPrefix("Hello world", "world"))
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, subffix string) bool {
	return len(s) >= len(subffix) && s[len(s)-len(subffix):] == subffix
}
