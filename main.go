package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {

	fmt.Println("Strings")

	var nameOne string = "first method for writing strings"
	var nameTwo = "Second method for writing strings"
	var nameThree string

	fmt.Println("Hello World, ", nameOne, ", ", nameTwo)

	nameThree = "Declared buh just used"
	nameOne = "updated"

	nameFour := "another method"
	fmt.Println("Hello World, ",nameOne, ", " ,nameTwo,  ", " ,nameThree,  ", ",nameFour)


	fmt.Println("Integers")

	var ageOne int = 29
	var ageTwo = 40
	ageThree:=27

	fmt.Println(ageOne, ageTwo, ageThree)

	fmt.Println("bits/memory")

	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255

	fmt.Println(numOne, numTwo, numThree)

	fmt.Println("float")
	var floatOne float32 = -1.5
	var floatTwo float64 = 567737832832823832.7
	floatThree:=23.6

	fmt.Println(floatOne, floatTwo, floatThree)

	fmt.Println(quote.Go())
}