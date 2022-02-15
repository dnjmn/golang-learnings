package main

import "fmt"

func main() {
    v := &[]int{}
    *v = append(*v, -1)
    fmt.Println(v)
}
