package main
import "fmt"

func main() {
character := "asSASA ddd dsjkdsjs dk"
for i := 0; i < len(character) ;i++ {

}

for index,value := range character{
	fmt.Printf("%#U  postion %s\n",value,index)
}
}