package main

import "fmt"

func main() {

	myArray := [10]int{12, 22, 24, 36}
	mySlice := myArray[:]
	fmt.Println("mySlice: ", mySlice)
	mySlice[0] = 99
	fmt.Println("slice: ", mySlice)

	myColors := [...]string{"Red", "Yellow", "Green", "Blue"}
	myColorSlice := myColors[1 : len(myColors)-1]
	fmt.Println("colorSlice: ", myColorSlice)

	nums := make([]int, 5)
	fmt.Println(nums)
	nums = append(nums, 112)
	fmt.Println(nums)

	fmt.Println("Length: ", len(nums))
	fmt.Println("Capacity: ", cap(nums))

	myArray2 := [10]int{2, 3, 5, 6, 7, 9, 12, 34, 36, 111}
	mySlice2 := myArray2[:]
	mySlice3 := myArray2[:]

	fmt.Println("O'zgarishsiz")
	fmt.Println("My array 2: ", myArray2)
	fmt.Println("My slice 2: ", mySlice2)
	fmt.Println("My slice 3: ", mySlice3)

	mySlice2[0] = 423

	fmt.Println("O'zgarish bilan")
	fmt.Println("My array 2: ", myArray2)
	fmt.Println("My slice 2: ", mySlice2)
	fmt.Println("My slice 3: ", mySlice3)

}
