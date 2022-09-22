package main

import "fmt"

func main() {
	count := 0
	fmt.Printf("iltimos sosn kiriting: ")
	fmt.Scanf("%d", &count)
	for i := 0; i < count; i++ {
		if i%2 == 0 {
			fmt.Printf(" Juft Son %d\n", i)
		}
	}

	from := 0 
	fmt.Printf("son kiriting:: ")
	fmt.Scanf("%d", &from)

	for i := from ; i > 0; i-- {
		fmt.Printf("Son: %d\n", i)
	}

}
