package main
import "fmt"

func main()  {
	i:=0
	start:
		if i<10 {
			fmt.Println(i)
			i += 1
			goto start			
		}else {
			goto End
		}
		End:
}