package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var grid [][]Tree

type Tree struct {
	Height            int
	VisibleFromLeft   bool
	VisibleFromRight  bool
	VisibleFromTop    bool
	VisibleFromBottom bool
}

func (t Tree) IsVisble() bool {
	if t.VisibleFromBottom == true || t.VisibleFromLeft == true || t.VisibleFromRight == true || t.VisibleFromTop == true {
		return true
	}
	return false
}
func main() {
	fname := "input.txt"
	//fname := "exampleinput.txt"

	readFile, err := os.Open(fname) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var row []Tree
		line := fileScanner.Text()
		for i := 0; i < len(line); i++ {
			treeheight, _ := strconv.Atoi(string(line[i]))

			row = append(row, Tree{Height: treeheight})
		}
		grid = append(grid, row)

	}
	visCount := 0
	//Part 1
	// Rows
	for r := 0; r < len(grid); r++ {
		maxFromL := -1
		maxFromR := -1
		maxFromT := -1
		maxFromB := -1
		numcols := len(grid[r]) - 1

		for c := 0; c <= numcols; c++ {
			// L to R
			if grid[r][c].Height > maxFromL {
				grid[r][c].VisibleFromLeft = true
				maxFromL = grid[r][c].Height
			}
			// R to L
			if grid[r][numcols-c].Height > maxFromR {
				maxFromR = grid[r][numcols-c].Height
				grid[r][numcols-c].VisibleFromRight = true
			}
			// And now fip the grid by 90 degrees
			// T to B
			if grid[c][r].Height > maxFromT {
				maxFromT = grid[c][r].Height
				grid[c][r].VisibleFromTop = true
			}
			// B to T
			if grid[numcols-c][r].Height > maxFromB {
				maxFromB = grid[numcols-c][r].Height
				grid[numcols-c][r].VisibleFromBottom = true
			
		}
	}
	for _, v := range grid {
		for _, v2 := range v {
			if v2.IsVisble() {
				visCount++
			}
		}
	}
	fmt.Println("Viscount", visCount)
	//Part 2
	
}
