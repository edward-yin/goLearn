package main

import (
	"fmt"
	// "os"
)

func main() {
	// fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("hello")

	var a [4]int = [4]int{5, 6, 7, 8}
	// for i := range a {
	// 	a[i] = i + 1
	// }

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
}
