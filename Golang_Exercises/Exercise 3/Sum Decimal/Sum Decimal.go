package main

import "fmt"
import	"math"
import	"math/big"
import	"strconv"


func add_sqrt(n int64) int {

	limit := new(big.Int).Exp(big.NewInt(10), big.NewInt(1001), nil)
	a := big.NewInt(5 * n)
	b := big.NewInt(5)
	five := big.NewInt(5)
	ten := big.NewInt(10)
	hundred := big.NewInt(100)

	for b.Cmp(limit) < 0 {
		if a.Cmp(b) < 0 {
			a.Mul(a, hundred)
			tmp := new(big.Int).Div(b, ten)
			tmp.Mul(tmp, hundred)
			b.Add(tmp, five)
		} else {
			a.Sub(a, b)
			b.Add(b, ten)
		}
	}
	b.Div(b, hundred)

	ansDec := b.String()

	sum := 0
	point := len(strconv.Itoa(int(math.Sqrt(float64(n)))))
	ansDec = ansDec[point:]
	for i := 0; i < len(ansDec); i++ {
		sum += int(ansDec[i] - '0')
	}
	return sum
}

func main() {
	fmt.Println(add_sqrt(3))
}
