package main
import "fmt"

func main() {
revers := "foobar"
for index := len(revers) -1;index >= 0; index--  {
	fmt.Printf("%s",string(revers[index]))
}
}