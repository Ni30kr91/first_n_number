package main

import "fmt"

func main() {

	fmt.Println("list of number from 1 to 100:")
	for i := 1; i <= 100; i++ {

		if i%2 == 0 {
			fmt.Print(" ", i)
		}
	}
	// fmt.Println()
	// for j := 1; j <= 100; j++ {
	// 	if j%j == 0 {
	// 		fmt.Print(" ", j)
	// 	}
	// }

}
