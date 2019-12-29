package main

import "fmt"

func mapfunc(sum func(x int) int, lst []int) []int {
	maps := []int{}
	for index := 0; index < len(lst); index++ {
		maps = append(maps, sum(lst[index]))
	}
	return maps
}

func main() {
	lst := []int{5, 7, 9, 13}
	f := func(x int) int {
		return x + x
	}
	fmt.Println(mapfunc(f, lst))
}