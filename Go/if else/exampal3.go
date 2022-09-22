package main

import "fmt"

func main() {
	// harbiy qabul
	// erkak 180 cm dan baland
	// ayol 160 cm dan baland

	var name, sex, height = "", "", 0

	fmt.Println("ismni kiriting: ")
	fmt.Scanln(&name)
	fmt.Println("jinsingizni kiriting")
	fmt.Scanln(&sex)
	fmt.Println("boyingizni kiriting")
	fmt.Scanln(&height)

	if sex == "erakak" && height > 180 {
		fmt.Printf("Salom %v. siz %v siz. boyingiz %v sm .\n", name, sex, height)
		fmt.Println("Qabul qilindiz")
	} else  if sex == "erkak" && height < 180 {
		fmt.Printf("Salom %v. siz %v siz. boyingiz %v sm .\n", name, sex, height)
		fmt.Println("Qabul qilinmadiz")
	}

}
