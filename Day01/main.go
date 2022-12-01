package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Part 1 - Read a file containing numbers in groups. Numbers are separated by CR, groups by a blank line terminated by CR
	// sum the numbers in each group and then find the group containing the highest number and output that number
	var calarray []int
	readFile, err := os.Open("Calories.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var individualTotal int
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			fmt.Println("BREAK")
			calarray = append(calarray, individualTotal)
			individualTotal = 0
		} else {
			linevalue, _ := strconv.Atoi(fileScanner.Text())
			individualTotal = individualTotal + linevalue
			fmt.Println("[", fileScanner.Text(), "]")

		}
	}
	var highestTotal = 0
	for _, v := range calarray {
		if v > highestTotal {
			highestTotal = v
		}
	}

	fmt.Println("The highest Total is :", highestTotal)

	// Now for part 2 = Take the summed groups and sum the three biggest groups to provide a total
	sort.Ints(calarray)
	var Top3Total int
	for i := len(calarray) - 1; i > len(calarray)-4; i-- {
		fmt.Println(calarray[i])
		Top3Total = Top3Total + calarray[i]

	}
	
	fmt.Println(Top3Total)

	readFile.Close()

}
