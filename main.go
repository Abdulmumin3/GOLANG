package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{1, 2, 3}
	var ages = [3]int{1, 2, 3}
	names:=[4]string{"Wale", "Taiwo", "Ade", "Shola"}
	names[1]="luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices (uses array under the hood)
	var scores = []int{23, 24, 26, 27}
	scores[2] = 25
	scores = append(scores, 28)

	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Tayo")
	fmt.Println(rangeOne)

}