package main

import (
    "fmt"
)

func reverse (check string) string{

	var toReverse [] rune
	var newString [] rune

	checkList:= []rune (check)
	var inPar = 0
	var start = 0
	var end = 0

	for i:=0; i<len(checkList);i++{

		if (checkList[i]=='('){
			if(start ==0){
				start = i
			}
			inPar++
		}else if (checkList[i]==')'){
			inPar--
			if(inPar==0){
			 	end = i
			}
		}else{
			if (inPar>0){
				toReverse = append(toReverse,checkList[i])
			}
		}	
	}

	

	for k,j:= 0,len(toReverse)-1 ; k<len(toReverse)/2; k,j= k+1,j-1 {
		toReverse[k],toReverse[j] = toReverse[j], toReverse[k]
	}

	newString = append(newString, checkList[:start]...)
	newString = append(newString, toReverse...)
	newString = append(newString,checkList[end+1:]...)

	return string(newString)

	


	

}

func main() {
	
	fmt.Println(reverse("foo(bar)baz"))
	fmt.Println(reverse("(bar)"))
	fmt.Println(reverse("foo(bar(baz))blim"))

}