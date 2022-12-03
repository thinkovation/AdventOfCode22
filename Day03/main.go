package main

/*

So for this one, I was lazy - there are only nine possible combinations. Simplest to create a map containing each combination and the associated score

*/

import (
	"bufio"
	"fmt"
	"os"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func letterScore(inletter byte) int {
	for k, v := range letters {
		if v == rune(inletter) {
			return k + 1

		}
	}
	return 0
}
func checkRucksack(input string) int {
	lettermap := make(map[string]bool)
	for _, v := range input[:len(input)/2] {

		lettermap[string(v)] = true

	}
	for _, v := range input[len(input)/2:] {
		if _, ok := lettermap[string(v)]; ok {
			return letterScore(byte(v))
		}
	}
	return 0
}
func lookForBadge(input []string) int {
	var itemcounts []map[string]int

	for _, v := range input {
		thismap := make(map[string]int)
		for _, item := range v {
			if _, ok := thismap[string(item)]; !ok {
				thismap[string(item)] = 1

			}
		}
		itemcounts = append(itemcounts, thismap)

	}
	for k, _ := range itemcounts[0] {
		if _, ok := itemcounts[1][k]; ok {
			if _, ok := itemcounts[2][k]; ok {
				return letterScore(byte(k[0]))
			}

		}
	}
	return 0
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var sumOfPriorities = 0
	for fileScanner.Scan() {
		// Do Something
		sumOfPriorities = sumOfPriorities + checkRucksack(fileScanner.Text())
	}
	fmt.Println("Part 1 :", sumOfPriorities)
	readFile.Close()
	readFile, err = os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner = bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	linecounter := 0
	sumOfBadges := 0
	var arraytosend []string
	for fileScanner.Scan() {
		// Do Something
		linecounter++
		arraytosend = append(arraytosend, fileScanner.Text())

		if linecounter == 3 {

			sumOfBadges = sumOfBadges + lookForBadge(arraytosend)
			linecounter = 0
			arraytosend = arraytosend[:0]
		}
	}
	readFile.Close()
	fmt.Println("Part 2 :", sumOfBadges)

}
