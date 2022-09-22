package main

import "fmt"

func main() {

	myArray := [...]int{1, 43, 5, 66, 90, 535, 99}

	fmt.Println(myArray[len(myArray)-1])

	myCars := [...]string{"lada", "porche", "mers"}

	fmt.Println(len(myCars))

	for i := 0; i < len(myArray); i++ {

		fmt.Println("number: ", myArray[i])

	}
}
