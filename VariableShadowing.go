package main 

import "fmt"

var a = 10 

func main() {
	age := 30

	if age > 18 {
		a := 25
		fmt.Println(a)
	}

	fmt.Println(a)
}