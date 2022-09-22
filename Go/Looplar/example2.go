package main

import (
	"fmt"
	
)

func main() {
	// sell ticket app
	countTicket := 100 
	isContunie := "y"
	for isContunie == "y" || isContunie == "Y" {
		fmt.Printf("bilet olishni istaysizmi[y,N]?: ")
		fmt.Scanf("%v", &isContunie)
		if isContunie != "y"  || isContunie == "Y"{
			fmt.Printf("Dastur tugatildi \n")
			break
		}
		ticketNum := 0
		fmt.Printf("nechta bilet sotib olasiz: ")
		fmt.Scanf("%d", &ticketNum)
		if ticketNum <= countTicket {
			countTicket -= ticketNum
			fmt.Printf("%d ta bilet sotib olindi \n", ticketNum)
			fmt.Printf("%d dona bilet qoldi sotuvda\n", countTicket)
		} else  if countTicket == 0{
		 fmt.Printf("biletlar tugadi dukon yopildi\n")
		} else {
			fmt.Println("Buncha bilet mavjud emas")
			fmt.Printf("%d dona bilet hozirda mavjud\n", countTicket)
		}
	
	}

}
