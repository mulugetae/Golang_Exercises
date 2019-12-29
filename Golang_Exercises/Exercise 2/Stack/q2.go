package main

import "fmt"

type stack2 struct {
	a1 [10]string
}

func (s *stack2) push(no string) {
	for i, val := range s.a1 {
		if i >= len(s.a1) && i >= 0 {
			return
		} else if val == "" {
			s.a1[i] = no
			return
		}
	}
}
func (s *stack2) pop() string {

	for i := len(s.a1) - 1; i >= 0; i-- {
		if s.a1[i] != "" {
			val := s.a1[i]
			s.a1[i] = ""
			return val
		}
	}
	return ""
}

func (s *stack2) printFormat() {
	fmt.Printf("My stack")
	for i, val := range s.a1 {
		fmt.Printf("[%v:%s]", i, val)

	}
	fmt.Println()
}

func main() {
	s := stack2{}
	fmt.Println(s)
	s.push("a")
	s.push("b")
	s.push("c")
	s.push("d")

	s.printFormat()
	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s)
	s.printFormat()
}