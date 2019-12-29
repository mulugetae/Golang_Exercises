package main

import (
	"fmt"
)
func main() {
	random, enc := encode(1990)
	fmt.Println(random, enc)
	deco, num := decode("MMVIII")
	fmt.Println(deco, num)
}

func encode(n int) (bool, string) {
	decimal := []int{1, 4, 5, 9, 10, 40,50, 90, 100, 400,500, 900, 1000}
	romrep := ""
	encoded := false
	roman := []string{
		"I", "IV", "V", "IX",
		"X", "XL", "L", "XC",
		"C", "CD", "D", "CM", "M"}
	for n > 0 {
		for i := len(decimal) - 1; i >= 0; i-- {
			if decimal[i] <= n {
				romrep += roman[i]
				n -= decimal[i]
			}
		}
	}

	if roman != nil {
		encoded = true
		return encoded, romrep
	} else {
		return encoded, romrep
	}

}
func decode(str string) (bool, int) {

	num := 0

	for i := 0; i < len(str); i++ {
		if i+1 > len(str) {
			if relateValue(string(str[i])) > relateValue(string(str[i+1])) {
				num += relateValue(string(str[i]))
			} else {
				num += relateValue(string(str[i+1])) - relateValue(string(str[i]))
			}
		} else {
			num += relateValue(string(str[i]))
		}

	}
	if num > 0 {
		return true, num
	} else {
		return false, num
	}

}

func relateValue(s string) int {
	decimal := []int{1, 4, 5, 9, 10, 40,
		50, 90, 100, 400,
		500, 900, 1000}
	roman := []string{
		"I", "IV", "V", "IX",
		"X", "XL", "L", "XC",
		"C", "CD", "D", "CM", "M"}
	val := 0
	for i := 0; i < len(roman); i++ {
		if roman[i] == s {
			val = decimal[i]
			break
		}
	}
	return val

}

