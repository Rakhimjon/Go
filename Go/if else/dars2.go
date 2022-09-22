package main

import "fmt"

func  main()  {
	fmt.Println("Tanlang: ")
	var tanlov int 
	fmt.Scanf("%d", &tanlov)

	switch tanlov  {
	case 1:
		fmt.Println("Balance 34$ ")
		break
	case 2: 
	fmt.Println("balance 4343 sum")
	break
	case 3: 
	fmt.Println("balance 98")
	break
	default:
		fmt.Println("Xato boldi")
		break
	}
}