package main

import "fmt"

func f(x,y int) (int,int){
	if y<x {
		return y,x
	}
	return x,y
}

func main() {
	fmt.Println(f(7,2))
	fmt.Println(f(2,7))
}