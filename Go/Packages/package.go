package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	
	fmt.Println("Hello gopher")

	fmt.Println("rondom son: ", rand.Intn(10))

	fmt.Println(strings.Contains("test", "es1"))
	fmt.Println(strings.HasPrefix("hellooooo", "hel"))
}