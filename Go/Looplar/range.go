package main
import "fmt"

var pow = []int{1,2,4,8,16, 32,64,128} 
var users = [] string {"Jhon", "Bernard","Asror"}
var cities = map[string]string{"Uz": "Tashkent","US": "Washington",} 
func main() {
	
	for i , value := range pow {
		fmt.Printf("2 -> %d = %d\n", i, value)
	}

	for _,v := range users {
		fmt.Printf("value: %s\n", v)
	}

	for key, value := range cities {
		fmt.Println("Map Key: ", key, "Value: ",value)
	}
}