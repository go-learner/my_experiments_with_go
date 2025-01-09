package main

import "fmt"

func main() {

	//var colors map[string]string

	//colors := make(map[string]string)

	//colors["red"] = "123345"
	colors := map[string]string{
		"red":    "#112345",
		"green":  "#23456",
		"yellow": "#987654",
	}
	printMap(colors)
}

func printMap(m map[string]string) {

	for color, hex := range m {
		fmt.Println("hex code for", color, "is", hex)

	}

}
