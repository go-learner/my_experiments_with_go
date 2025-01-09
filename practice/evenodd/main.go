package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range num {
		if i%2 == 0 {
			fmt.Printf("the number is %v which is even\n", i)
		} else {
			fmt.Printf("the number is %v which is odd\n", i)
		}
	}
	fmt.Println(num)

}
