package main 

import "fmt"

func main()  {
	
	//baholash tizimi 
	var rate int 
	fmt.Println("balni kiriting: ")
	fmt.Scanf("%d", &rate)
	if rate > 100 {
		fmt.Println("grant :::")
	} else if rate >= 80 && rate >= 100 {
		fmt.Println("5 baho")
	} else if rate > 60 && rate < 80 {
		fmt.Println("4 baho")
	}  else if rate > 40 && rate < 60 {
		fmt.Println("3 baho")
	} else {
		fmt.Println("failed")
	}

	// || OR yoki -> +
	fmt.Println(true || true && false) //1+1*0 = 1
	fmt.Println(true || false ) //1+0 = 1
	fmt.Println(false || true ) //0+1 = 1
	fmt.Println(false || false ) //0+0 = 0


	// && AND va  -> *
	fmt.Println(true && true ) //1 * 1  = 1
	fmt.Println(true && false )// 1*0 = 0
	fmt.Println(false && true )// 0*1 = 0
	fmt.Println(false && false )// 0*0 = 0
}