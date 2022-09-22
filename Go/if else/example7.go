package main

import "fmt"

func main() {
	//krirtilgan harfni katta yoki kichik ekanligini aniqlash
	// ASCII
	harf := 't'

	if harf <= 90 &&  harf >= 67  {
		fmt.Println("kiritilgan harf katta")
	} else if harf <= 122 && harf >= 97 {
		fmt.Println("kritilgan har kichkina")
	}
}