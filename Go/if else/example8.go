package main

import "fmt"

func main() {
	
	 asciiValue := 0
	fmt.Println("Son kiriting")
	fmt.Scanln(&asciiValue)
	character := rune(asciiValue)

	fmt.Printf("Ascii desimal %d => %c\n", asciiValue, character)
}