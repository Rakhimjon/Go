package main
import "time"

func main() {
	
	k := 0
	summ := 0 

	for k < 10 {
		println("k:", k)
		summ += k
		k += 1
		time.Sleep(500 * time.Millisecond)
	}

	print("yigindi: ", summ)
}