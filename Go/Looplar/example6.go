package main
import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d x %d = %d\t ", j, i, j*i)
		}
		fmt.Println("\n")
	}
}