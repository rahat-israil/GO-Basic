package main

import "fmt"

var (
	a = 10
	b = 20
)

func add (x int, y int){
	z := x + y
	printNum(z)
}
func printNum(num int){
	fmt.Println(num)
}
func main (){
	add(a, b)
}