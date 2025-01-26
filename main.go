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

func main4 (){
    a := 20
    b := 10

    p, q := getNumbers (a, b);

    fmt.Println(p)
    fmt.Println(q)
}


func main5 (){
    // Print welcome message
    fmt.Println("Welcome to the Calculator System");

    // Input Fields
    var name string;
    fmt.Print("Please Enter your name: ");
    fmt.Scanln(&name);

    fmt.Println("Here you can calculate the 2 number Summation, Subtraction, Multiplication and Division");

    var num1 int;
    var num2 int;

    fmt.Print("Please Enter thr first number: ")
    fmt.Scanln(&num1);

    fmt.Print("Please Enter the Second Number: ");
    fmt.Scanln(&num2);

    // Operations
    sum := num1 + num2;
    sub := num1 - num2;
    multi := num1 * num2;
    divide := num1 / num2;

    // Final Print
    fmt.Println("Hello!", name);
    fmt.Println("The sum of the number is:", sum);
    fmt.Println("The Subtrsct of the number is:", sub);
    fmt.Println("The Multiplication of the number is:", multi);
    fmt.Println("The Division of the number is:", divide);

    fmt.Println("Thank you for using the Calculator System");
}

// Organized way to implement the Function

func printWelcomeMessage (){
    fmt.Println("Welcome to the Calculator System");
}
func getUserName () string{
    var name string;
    fmt.Print("Please Enter your name: ");
    fmt.Scanln(&name);

    return name;
}
func getTwoNumbers () (int, int){
    fmt.Println("Here you can calculate the 2 number Summation, Subtraction, Multiplication and Division");
    var num1 int;
    var num2 int;

    fmt.Print("Please Enter thr first number: ")
    fmt.Scanln(&num1);

    fmt.Print("Please Enter the Second Number: ");
    fmt.Scanln(&num2);

    return num1, num2;
}
func Operations(num1 int, num2 int) (int, int, int, int) {
    sum := num1 + num2;
    sub := num1 - num2;
    multi := num1 * num2;
    divide := num1 / num2;

    return sum, sub, multi, divide;
}
func DisplayResult(name string, sum int, sub int, multi int, divide int){
    fmt.Println("Hello!", name);
    fmt.Println("The sum of the number is:", sum);
    fmt.Println("The Subtrsct of the number is:", sub);
    fmt.Println("The Multiplication of the number is:", multi);
    fmt.Println("The Division of the number is:", divide);
}
func printGoodByeMessage(){
    fmt.Println("Thank you for using the Calculator System");
    fmt.Println("Good Bye!");
}

func main(){
    printWelcomeMessage();
    name := getUserName();
    num1, num2 := getTwoNumbers();

    sum, sub, multi, divide := Operations(num1, num2);

    DisplayResult(name, sum, sub, multi, divide);
    printGoodByeMessage();
}