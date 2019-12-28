package main
import "fmt"

func main() {
for index := 0; index < 100; index++ {
	for index1 := 0; index1 < index; index1++ {
		fmt.Print("A")
	}
	fmt.Println()
}
}