package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	resp, err := http.Get("http://standards-oui.ieee.org/oui/oui.txt")
	if err != nil {
		fmt.Println("unable to fetch the url", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: status code %d\n", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := string(body)
	lines := strings.Split(s, "\n")

	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := writer.Write(lines)
	fmt.Println(record)

}
