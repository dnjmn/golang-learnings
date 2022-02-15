package main

import (
	"fmt"
	"time"
)

func main() {}

func run(f bool) int64 {
	i := 0
    v2 := ""
    t1 := time.Now().UnixNano()
	for i < 100000000000 {
        if f {
		    v1 := "x"
		    _ = v1
        } else {
            v2 = "x"
		    _ = v2
        }
        i++
	}
    t2 := time.Now().UnixNano()
    return (t2-t1)
}
