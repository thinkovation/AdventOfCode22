package main

import (
	"fmt"
	"os"
)

func isMixed(in string) bool {

	letterMap := make(map[string]int)
	for i := 0; i < len(in); i++ {
		if _, ok := letterMap[string(in[i])]; ok {
			return false
		} else {
			letterMap[string(in[i])] = 0

		}

	}
	return true
}

// Part 1
func scanString(in string) int {

	for i := 4; i < len(in); i++ {
		if isMixed(in[i-4 : i]) {

			return i
		}
	}
	return 0
}

// Part 2 (Could have been a parameter to the original scanString but to preserve the part1 / part 2 work I copied part 1 and changed the length of the window to check)
func scanStringForMessageMarker(in string) int {
	fmt.Println("Len String", len(in))
	for i := 14; i < len(in); i++ {
		if isMixed(in[i-14 : i]) {

			return i
		}
	}
	return 0
}

func main() {
	fname := "input.txt"
	//fname := "exampleinput.txt"

	b, err := os.ReadFile(fname) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	firstMarker := scanString(str)
	fmt.Println("First Marker :", firstMarker)
	firstMessage := scanStringForMessageMarker(str)
	fmt.Println("First Message:", firstMessage)
}
