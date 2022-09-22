package main

import (
	"fmt"
	
)

func main() {
	states := make(map[string]string)
	
	states["TAS"] = "Tashkent"
	states["NAM"] = "Namangan"
	states["FER"] = "Fergana"
	delete(states, "NAM")
	fmt.Println("states: ", states)

	fmt.Println("state one: ", states["NAM"])

	for k, v := range states {
		fmt.Printf("Key: %v. value: %v\n", k,v)
		
	}
	
	myMap := make(map[int]string)
	myMap[12] = "jkj"
	myMap[32] = "kkk"

	keys := make([] string, len(states))

	i := 0

	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println("Keys slice: ", keys)
}