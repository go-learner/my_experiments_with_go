package main

import "fmt"

func main() {

	s := []string{"hi", "how", "are", "you"}
	updateSlice(s)
	fmt.Println(s)
}

func updateSlice(a []string) {
	a[0] = "bye"
}
