package main

import "fmt"

func main() {
	
	list := []int{4, 2, 3} 

	fmt.Println(missing(list))

}


func missing (nums []int) []int{

	var miss [] int
	var found = false

	var length = len(nums)
	length = length+2


	for index:=1; index<=length;index++{

		for i:=0; i<len(nums);i++{
			if (index==nums[i]){
				found = true
				break
			}
		}

		if (!found){
			miss = append(miss,index)
		}
		found = false
		
	}

	return miss


}