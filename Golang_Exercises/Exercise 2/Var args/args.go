package main

import "fmt"
func main() {
	printArgs(2,3)
	printArgs(2,3,5,6)
}
func printArgs (y ...int){

	for _,y :=  range y{
		fmt.Println(y)
	}

}