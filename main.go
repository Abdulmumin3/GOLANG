package main

import "fmt"

func main() {
	age:=24
	name:="Wale"

	//Print
	fmt.Print("Hello, ")
	fmt.Print("World \n")
	fmt.Print("New line \n")

	//Println
	fmt.Println("Hello, World")
	fmt.Println("New Line")

	fmt.Println("my age is", age, "and my name is", name)

	//Printfn(formatted strings) %_ = format specifiers
	fmt.Printf("my age is %v, and my name is %v \n", age, name)
	fmt.Printf("my age is %q, and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points \n", 225.55)
	fmt.Printf("you scored %0.1f points \n", 225.55)

	//Sprint (save formatted strings)
	var str = fmt.Sprintf("my age is %v, and my name is %v \n", age, name)
	fmt.Printf("The saved string is: %v", str)
}