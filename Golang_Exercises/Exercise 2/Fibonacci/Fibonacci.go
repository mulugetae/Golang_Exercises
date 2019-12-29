package main

import "fmt"
func fibo(x int) []int {
	fibonachi := []int{}
	sum := 1
	fibonachi = append(fibonachi, sum)
	for index := 1; index < x; index++ {
		if index < 2 {
			fibonachi = append(fibonachi, index)
		} else {
			sum = fibonachi[index-1] + fibonachi[index-2]
			fibonachi = append(fibonachi, sum)
		}
	}
	return fibonachi
}
func main()  {
	fmt.Println( fibo(7))
}

