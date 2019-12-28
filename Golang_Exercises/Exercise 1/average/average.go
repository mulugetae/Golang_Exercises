package main

import "fmt"

func main() {
	floatnumber := []float64{6.0,5.0,9.0}
	sum := 0.0
	for i :=0; i <len(floatnumber); i++ {
		sum += floatnumber[i]
	}
	average := sum/float64(len(floatnumber))
	fmt.Println("the avarage of the value is ", average)
}