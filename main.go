package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "hello there friends"
	// fmt.Println(strings.Contains(greeting, "ther"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "h"))
	// fmt.Println(strings.Split(greeting, " "))


	// fmt.Println("the original sentence = ",greeting)
	// fmt.Printf("the original sentence = %v", greeting)

	ages:=[]int{23, 24, 25, 26, 27, 28}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 26)
	fmt.Println(index)

	names:=[]string{"Wale", "Ade", "Tola", "Shayo"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Ade"))
}