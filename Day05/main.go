package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Stacks [10][]string
var StackIdx int

func addRowToStacks(inline string) {
	//1,5,9,13,17,21,25,33
	for i := 1; i < len(inline); i = i + 4 {
		//fmt.Println(string(inline[i]))
		if string(inline[i]) != " " {
			StackIdx := ((i - 1) / 4) + 1

			Stacks[StackIdx] = append([]string{string(inline[i])}, Stacks[StackIdx]...)
		}
	}
}
func move(n, f, t int) {
	for i := len(Stacks[f]) - 1; i >= len(Stacks[f])-n; i-- {
		Stacks[t] = append(Stacks[t], string(Stacks[f][i]))
	}
	Stacks[f] = Stacks[f][:len(Stacks[f])-n]
}
func moveFast(n, f, t int) {
	Stacks[t] = append(Stacks[t], Stacks[f][len(Stacks[f])-n:]...)
	Stacks[f] = Stacks[f][:len(Stacks[f])-n]
}

func main() {
	readFile, err := os.Open("input.txt")
	//readFile, err := os.Open("exampleinput.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var loadingStacks bool = true
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			loadingStacks = false

		}

		if loadingStacks {
			if string(fileScanner.Text()[1]) != "1" {
				addRowToStacks(fileScanner.Text())
			}
		} else {

			commandarray := strings.Split(fileScanner.Text(), " ")
			if len(commandarray) > 4 {

				from, _ := strconv.Atoi(string(commandarray[3]))
				num, _ := strconv.Atoi(string(commandarray[1]))
				to, _ := strconv.Atoi(string(commandarray[5]))
				// Part 1
				// move((num, from, to))
				// Part 2
				moveFast(num, from, to)

			}
		}

	}

	var topsofstacks string

	for i := 1; i < len(Stacks); i++ {
		if len(Stacks[i]) > 0 {
			topsofstacks = topsofstacks + Stacks[i][len(Stacks[i])-1]
		}
	}
	fmt.Println(topsofstacks)
	readFile.Close()

}
