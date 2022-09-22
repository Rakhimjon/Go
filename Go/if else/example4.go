package main
import "fmt"
func main() {
	// chakana savdo dukoni 
	// apelsin 
	// 1 -10 -> 3$
	// 11 -100 -> 2.75$
	// 101 - 1000 -> 2.5$
	//1001 - 10000 -> 	2$ 

	//1001.... -> 1.25 $

	var dona , donaNarx  = 0, 0.0 

	fmt.Println("Necha dona apelsin olishni istaysiz")
	fmt.Scanf("%d", &dona)

	if dona > 10000 {
		donaNarx = 1.25
	} else if dona > 1000 && dona <= 10000 {
		donaNarx = 2.0
	} else if dona > 100 && dona <= 1000 {
		donaNarx = 2.5
	} else if dona > 10 && dona <=  100 {
		donaNarx = 2.75
	} else if dona > 0 && dona <= 10 {
		donaNarx = 3
	} else {
		fmt.Println("faqat musbat son kiriting")
	}
	var jami float32
	jami = float32(dona) * float32(donaNarx)

	fmt.Println("jami tulov sum", jami)


}