package main

import "fmt"

func add(x int) func(num int) int {

	fnc := func(num int) int {
		return num + x
	}
		return fnc
}

func main() {
	anotherFUN_call := add(6)
	fmt.Println(anotherFUN_call(2))
}
