package main

import "fmt"

func main() {

	son := 0
	fmt.Printf("Iltimos son kiriting: ")
	fmt.Scanf("%d", &son)

	yigindi := 0
	factorial := 0

	for i := 1; i <= son; i++ {
		yigindi += i
		factorial *= i
	}
	fmt.Printf("yigindi: %d\n", yigindi)
	fmt.Printf("yigindi: %d\n", factorial)
}
