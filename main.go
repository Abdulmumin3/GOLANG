package main

import (
	"fmt"
)


func main() {
	myBill:=newBill("Wale's bill")

	fmt.Println(myBill.format())
}