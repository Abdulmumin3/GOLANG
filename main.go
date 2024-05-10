package main

import (
	"fmt"
)

func main() {
	// age:=50
	// fmt.Println(age>=45)
	// fmt.Println(age<=50)
	// fmt.Println(age==45)
	// fmt.Println(age!=50)

	// if age<35{
	// 	fmt.Println("age is less than 35")
	// } else if age<50{
	// 	fmt.Println("age is less than 50")
	// } else{
	// 	fmt.Println("age is not less than", age)
	// }

	names:=[]string{"Ade", "Shola", "Shade", "Tayo", "bayo"}
	
	for index, value := range names {
		if index == 1{
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2{
			fmt.Println("breaking at pos", index)
			break
		}

		fmt.Printf("the value at pos %v is %v \n",index, value)
	}
}