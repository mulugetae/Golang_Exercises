package main

import "fmt"

func main() {
	list := []int{4, 3, 2, 6, 1}
	fmt.Println(buble_Sort(list))
}
func buble_Sort(bublist []int) []int {
	leng := len(bublist)
	for i := 0; i < leng; i++ {
		//fmt.Println(bub)
		for index := 1; index < leng; index++ {
			if bublist[index-1] > bublist[index] {
				bublist[index-1], bublist[index] = bublist[index], bublist[index-1]

			}
		}
		fmt.Println(bublist)

	}
	fmt.Println("this is sorted list \n")
	return bublist
}