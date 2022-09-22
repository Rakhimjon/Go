package main

import "fmt"

func main() {
	// juft son yoki toq son aniqlash
	son := 0

	fmt.Println("juft son yoki toq son kiriting: ")

	fmt.Scanf("%d", &son)

	if son%2 == 0 {
		fmt.Println(son, "Juft  : ")
	} else {
		fmt.Println(son, "toq")
	}
}
