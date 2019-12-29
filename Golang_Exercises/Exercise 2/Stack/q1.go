package main

import "fmt"

type stack struct {
	value []string
}

func (s *stack) push(str string){
	s.value = append(s.value, str)
}
func (s *stack) pop()  {
	s.value = s.value[:len(s.value)-1]

}

func main() {
	 value :=[]string{"1","2"}
	 s:=stack{value:value}
	 s.push("add")
	 fmt.Println(s.value)
	 s.pop()
	 fmt.Println(s.value)
}