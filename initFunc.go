package main 

import "fmt"

var a = 20

func main (){
	fmt.Println(a)
}
func init(){
	fmt.Println(a)
	a = 30
}