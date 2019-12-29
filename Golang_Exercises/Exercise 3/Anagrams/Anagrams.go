package main

import "fmt"
import "strings"

func anagram (checks []string, ana string) []string{

	var anagrams [] string

	for k:=0; k<len(checks);k++{
		ana = strings.ToLower(ana)
		var check = checks[k]
		check = strings.ToLower(check)
		
		checkList:= []rune (check)
		anaList :=[]rune (ana)

		for i := 0; i<len(anaList); i++{
			for j := 0; j<len(checkList); j++{
				if (anaList[i]==checkList[j]){
					checkList = append(checkList[:j], checkList[j+1:]...)
					break
				}

			}
		}

		if (len(checkList)==0){
			anagrams = append(anagrams,check)
		}
	}

	return anagrams

	

}

func main() {
	checks_words := []string{ "Madaj Curie","Madam Curie"} 
	fmt.Println(anagram(checks_words,"Radium came"))
}