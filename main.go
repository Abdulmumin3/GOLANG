package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func getInput (prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill{
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Creating the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose Option: (a - add item), (s - save bill), (t - add tip) ", reader)	
	fmt.Println(opt)
}


func main() {
	// myBill:=newBill("Wale's bill")

	// myBill.addItem("onion soup", 4.50)
	// myBill.addItem("veg pie", 8.95)
	// myBill.addItem("toffee pudding", 4.95)
	// myBill.addItem("coffee", 3.25)

	// myBill.updateTip(10)

	// fmt.Println(myBill.format())

	myBill := createBill()
	promptOptions(myBill)
	// fmt.Println(myBill)
}