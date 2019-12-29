package main

import "fmt"

func buildWord(word string,dict []string) int {
    res := 0
    done := false
	for i:=0;i<len(dict) ; i++ {
		sen := ""
		num := 0
		for j :=i; j<len(dict) ;j++  {
			sen += dict[j]
			num++
			if sen == word{
				done = true
				break
			}
		}
		if done{
			res = num
			break
		}
	}

	return res

}

func main() {
	fmt.Println(buildWord("abracadabra", []string{"ab", "ra", "cadabra", "cada", "bra", "abracadabra"}))
	fmt.Println(buildWord("build", []string{"b", "u", "il", "ld"}))
	
}