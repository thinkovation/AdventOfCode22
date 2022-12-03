package main

/*

So for this one, I was lazy - there are only nine possible combinations. Simplest to create a map containing each combination and the associated score

*/

import (
	"bufio"
	"fmt"
	"os"
)

func CalculateScoreRevised(inline string) int {
	var scoretable = map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}

	return scoretable[inline]
}

func CalculateScore(inline string) int {
	var scoretable = map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	return scoretable[inline]
}

func main() {
	// Part1
	var runningscore int
	// Part2
	var runningscorepart2 int
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		// Part 1
		runningscore = runningscore + CalculateScore(fileScanner.Text())
		// Part 2
		runningscorepart2 = runningscorepart2 + CalculateScoreRevised(fileScanner.Text())
	}

	fmt.Println("The score is :", runningscore)
	fmt.Println("The part 2 score is :", runningscorepart2)

}
