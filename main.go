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

func main() {
    y := 10;

    if (y > 10 && y < 20) {
        fmt.Println("y is greater than 10");
    } else if (y > 9 || y == 20) {
        fmt.Println("Special case")
    } else {
        fmt.Println("Nothing");
    }
}