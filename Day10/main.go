package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var CycleCounter int
var XVal int
var XValHistory []int
var Vpos int
var Hpos int

func Tick() {
	CycleCounter++
	XValHistory = append(XValHistory, XVal)
	PaintPixel()

}

var display [6][40]string

func PaintPixel() {

	if Hpos > 39 {
		Hpos = 0
		Vpos++
		if Vpos > 5 {
			Vpos = 0
		}
	}

	if Hpos > XVal-2 && Hpos < XVal+2 {
		display[Vpos][Hpos] = "#"
	} else {
		display[Vpos][Hpos] = "."
	}
	Hpos++
}
func CalSignalStrength(cycles []int) int {
	var currentTotal int
	for _, v := range cycles {
		currentTotal = currentTotal + (XValHistory[v] * v)
	}
	return currentTotal
}
func main() {
	fname := "input.txt"
	//fname := "exampleinput.txt"
	// Put 0 value into history array to fill up the 0th index
	XValHistory = append(XValHistory, 0)
	XVal = 1

	readFile, err := os.Open(fname) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		instruction := fileScanner.Text()
		if instruction == "noop" {
			Tick()

		} else {
			cmdandparams := strings.Split(instruction, " ")
			if cmdandparams[0] == "addx" {
				numtoadd, _ := strconv.Atoi(cmdandparams[1])
				Tick()
				Tick()
				XVal = XVal + numtoadd
			}
		}
		fmt.Println(fileScanner.Text())

	}
	fmt.Println(XValHistory[1:])
	fmt.Println(len(XValHistory[1:]))
	fmt.Println(CalSignalStrength([]int{20, 60, 100, 140, 180, 220}))

	for _, v := range display {
		row := ""
		for _, v2 := range v {
			row = row + v2
		}
		fmt.Println(row)
	}
	fmt.Println("Done ")
	// ELPLZGZL
}
