package main

import "fmt"

func main() {
	x, y, z := 1, 1, 0
	for i := 0; i < 10; i++ {
		z = x + y
		x = y
		y = z

		fmt.Printf("%d ", z)
	}
}
