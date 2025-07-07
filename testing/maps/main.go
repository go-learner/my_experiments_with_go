package main

import "fmt"

var mapData map[string]string

func init() {
	fmt.Println("init calling")
	mapData = make(map[string]string)
}

func main() {
	fmt.Println("this is testing directory")
	addItem("ace of", "hearts")
	upDateItem("queen of", "diamonds")
	getById("ace of")
	deleteItem("queen of")
}

func addItem(k string, v string) {
	mapData[k] = v
	fmt.Println(mapData)
}

func upDateItem(k, v string) {
	mapData[k] = v
	fmt.Println("update map value:", mapData)
}

func getById(k string) {
	fmt.Println(mapData[k])
}

func deleteItem(k string) {
	delete(mapData, k)
	fmt.Println(mapData)
}
