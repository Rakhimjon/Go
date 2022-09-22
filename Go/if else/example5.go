package main
import "fmt"

func main() {
	// birga qoldiqsiz bolinishi va bolinmasligi 

	kattaSon, kichikSon := 0,0 

	fmt.Println("katta son: ")
	fmt.Scanf("%d",&kattaSon)
	fmt.Println("kichik son")
	fmt.Scanf("%d", &kichikSon)

	if kattaSon%kichikSon == 0 {
		fmt.Println("bolinadi ")
	}else {
		fmt.Println("bolinmaydi")
	}
}