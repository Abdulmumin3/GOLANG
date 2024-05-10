package main

import (
	"fmt"
)

func main() {

	menu:=map[string]float64{
		"soup": 789.6,
		"garri": 678.9,
		"epa": 456.7,
	}

	fmt.Println(menu)
	fmt.Println(menu["epa"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phoneBook:=map[int]string{
		8020736247: "Wale",
		9078965756: "Shade",
		9197654567: "Tayo",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[8020736247])

}