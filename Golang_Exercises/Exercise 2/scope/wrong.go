package main

import "fmt"

func main() {
	for i:=0; i < 10; i++{
		fmt.Printf("%v\n",i)
	}
	
// fmt.Printf("%v\n",i) //i is out of scpe it decleraed inside the for block and undefined: i
}