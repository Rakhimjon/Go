package main

import "fmt"

func main() {
	// siz kiritgan sondan 0 gacha bolgan sonlar orasidan 4 ga bo'linadigan
	// larni ekranga chiqarish 
	son := 0
	fmt.Scanf("%d", &son)

	for ; son > 0; son-- {
		if(son % 4 == 0){
			fmt.Printf("Bòlinadi: %d\n", son)
		}  else {
			
		}
	}

}