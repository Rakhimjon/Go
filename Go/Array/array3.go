package main

import "fmt"

func main() {

	contries := [...]string{"UK", "Uzbekistan", "Pakistan", "US"}

	for i, v := range contries {

		fmt.Println("index: ", i, "value: ", v)

	}

	numbers := [...]int{3, 2, 4, 6, 7, 8, 54, 545, 66}

	for _, v := range numbers {
		if v%2 == 0 {
			fmt.Println("juft son: ", v)
		}
	}

}
