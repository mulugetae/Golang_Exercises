package main

import "fmt"

func Min(slice []int) int {
	min := slice[4]
	for index := 1; index < len(slice); index++ {
		if (min == index){
			break
		}else{
			if min < slice[index] {
				min = slice[index]
			}
		}
		
	}
	return min
}

func main() {
	slice := []int{6, 15, 3, 7, 10}
	fmt.Println(Min(slice))

}