package main

import "fmt"

func main() {
	caller := new(Hello)
	caller.print()
	caller1 := new(Hello1)

	caller1.print()
}

type Hello struct {
	toPrint string
}

func (h Hello) print() {
	fmt.Println("Hello Printing", h.toPrint)
}

type Hello1 Hello

func (h Hello1) print() {
	fmt.Println("Hello Printing", h.toPrint)
}
