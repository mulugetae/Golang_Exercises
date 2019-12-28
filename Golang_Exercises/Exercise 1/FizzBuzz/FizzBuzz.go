package main
import "fmt"


func main() {
for index := 1; index < 100; index++ {
	if  index % 5== 0 && index % 3== 0{
		fmt.Println("Fizz Buzz",index)
	}else if index % 3== 0 {
		fmt.Println("Fizz",index)
	}else if  index % 5== 0 {
		fmt.Println("Buzz",index)
	}
}

}