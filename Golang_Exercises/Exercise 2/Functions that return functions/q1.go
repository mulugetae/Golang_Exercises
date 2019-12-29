package main

import "fmt"

func add() func(num int) int {

	fnc := func(num int) int {
		return num + 2
	}
		return fnc
}

func main() {
	anotherFUN_call := add()
	fmt.Println(anotherFUN_call(2))
}