package main
import (
	"fmt"
)

type toreset []byte

func main(){
	var sinitial,sresult toreset
	var c = "input string here"
	fmt.Scan(&c)
	sinitial,sresult = []byte(c),[]byte(c)
	order := make([]int,len(sinitial))
	for i :=0;i < len(sinitial);i++ {
		fmt.Scan(&order[i])
	}
	for j,v := range order {
		sresult[v] = sinitial[j]
	}
	fmt.Printf("%c",sresult)
}

