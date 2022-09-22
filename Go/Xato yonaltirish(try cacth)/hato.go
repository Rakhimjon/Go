package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Fayl: ", file.Name())
	}

	myErrors := errors.New("my error")
	println("Err: ", myErrors)
}

func chekError(err error) {
	if err != nil {
		println("error: ", err.Error())

		os.Exit(1)
	}
}
