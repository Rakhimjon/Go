package main

import (
	"fmt"
	"runtime"
	"time"
)

func alphabet() {
	for l := byte('a'); l <= byte('k'); l += 1  {
		 go fmt.Println(string(l))
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go alphabet()
	time.Sleep(time.Second)
}
