package main

import "fmt"


// func main() {
// 	fmt.Println("Hello, World!");
// }


// Variables & Data Types
// func main() {
//     var name = "Rahat";
//     age := 24;
//     var x int = 5;

//     fmt.Println(name, age, x);
// }

//  Function
func main1(){
    a := 5;
    b := 6;

    sum := a * b;

    fmt.Println(sum);
}


// if else
func main2() {
    x := 5;

    if (x > 10) {
        fmt.Println("x is greater than 10");
    } else {
        fmt.Println("x is less than 10");
    }
}

func main3() {
    y := 10;

    if (y > 10 && y < 20) {
        fmt.Println("y is greater than 10");
    } else if (y > 9 || y == 20) {
        fmt.Println("Special case")
    } else {
        fmt.Println("Nothing");
    }
}

// Function
func add (num1 int, num2 int){
    sum := num1 + num2;
    fmt.Println(sum);
}

func multiply (num1 int, num2 int) int {
    multi := num1 * num2;
    return multi;
}
// func main () {
//      a := 5;
//      b := 10;

//     result := multiply(a, b);  // 15
    

//     fmt.Println(result);
// }


func getNumbers (num1 int, num2 int) (int, int){
    sum := num1 + num2;

    divide := num1 / num2;

    return sum, divide;
}

func main (){
    a := 20
    b := 10

    p, q := getNumbers (a, b);

    fmt.Println(p)
    fmt.Println(q)
}