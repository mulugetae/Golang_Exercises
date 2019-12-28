package main
import  "fmt"
func main()  {
	//intalize the array
	x := []int{}
	for index := 0; index < 10; index++ {
		//insert the number
		x= append(x,index)
	}
	fmt.Println(x)
}

