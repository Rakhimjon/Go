package main
//OUTPUT - Input
import ("fmt"
		"bufio"
		"os"
		"strconv"
	)

func main()  {

	println("Hello go")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Son kiritung")

	son,_ := reader.ReadString('\n')

	fmt.Print(son)

	fmt.Println("Number kiriting")
	str, _ := reader.ReadString('\n')
	value,_ := strconv.ParseFloat(str, 64)
	fmt.Println(value +  100)
 

}