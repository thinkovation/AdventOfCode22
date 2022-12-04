package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func oneContainsTheOther(instring string) bool {
	parts := strings.Split(instring, ",")
	aparts := strings.Split(parts[0], "-")
	bparts := strings.Split(parts[1], "-")
	amin, _ := strconv.Atoi(aparts[0])
	amax, _ := strconv.Atoi(aparts[1])
	bmin, _ := strconv.Atoi(bparts[0])
	bmax, _ := strconv.Atoi(bparts[1])

	if amin <= bmin && amax >= bmax {
		return true
	}
	if bmin <= amin && bmax >= amax {
		return true
	}
	return false
}
func oneOverlapsTheOther(instring string) bool {
	parts := strings.Split(instring, ",")
	aparts := strings.Split(parts[0], "-")
	bparts := strings.Split(parts[1], "-")
	amin, _ := strconv.Atoi(aparts[0])
	amax, _ := strconv.Atoi(aparts[1])
	bmin, _ := strconv.Atoi(bparts[0])
	bmax, _ := strconv.Atoi(bparts[1])

	if amin >= bmin && amin <= bmax || amax >= bmin && amax <= bmax || bmax >= amin && bmax <= amax || bmin >= amin && bmin <= amax {
		return true
	}

	return false
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	containedcounter := 0
	overlapscounter := 0
	for fileScanner.Scan() {

		if oneContainsTheOther(fileScanner.Text()) {

			containedcounter++
		}
		if oneOverlapsTheOther(fileScanner.Text()) {

			overlapscounter++
		}

	}
	fmt.Println("Contained :", containedcounter)
	fmt.Println("Overlapped :", overlapscounter)
	readFile.Close()

}
