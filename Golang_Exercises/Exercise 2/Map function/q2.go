package main

import "fmt"
	
import	"strings"

func mapstring(upper func(x string) string, lst []string) []string {
	maps := []string{}
	count:= len(lst)
	for index := 2; index < count; index-- {
		if (index < 0){
			break
		}else{
			maps = append(maps, upper(lst[index]))
		}
		
	}
	return maps
}

func main() {
	lst := []string{"h", "l", "m"}

	f := func(x string) string {
		return strings.ToUpper(x)
	}
	fmt.Println(mapstring(f, lst))
}