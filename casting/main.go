package main

import "fmt"

func main() {
	i := get()
	fmt.Println(int(i))
}

func get() interface{} {
	return "hel"
}
