package main

import "fmt"

func main() {
	s := ", hello"
	s = s[2:]
	fmt.Println(len(s))
}
