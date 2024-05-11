package main

import (
	"fmt"
)

func updateName(x string) string{
	x="wedge"
	return x
}

func updateMenu(y map[string]float64){
	y["coffee"]=2.99
}

func main() {

	name:="tifa"

	name=updateName(name)

	fmt.Println(name)

	menu:=map[string]float64{
		"soup": 789.6,
		"garri": 678.9,
		"epa": 456.7,
	}

	updateMenu(menu)
	fmt.Println(menu)
}