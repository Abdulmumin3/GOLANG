package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string)  {
	fmt.Println("Good morning",n)
}

func sayBye(n string){
	fmt.Println("Good bye",n)
}

func cycleNames(n []string, f func(string)){
	for _, v:= range n{
		f(v)
	}
}

func circleArea(r float64) float64{
	return math.Pi * r * r
}

func main() {
	// sayGreeting("Wale")
	// sayBye("Shola")

	cycleNames([]string{"Wale", "Shola", "Bayo"}, sayGreeting)
	cycleNames([]string{"Wale", "Shola", "Bayo"}, sayBye)

	a1:=circleArea(23.4)
	a2:=circleArea(12.0)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}