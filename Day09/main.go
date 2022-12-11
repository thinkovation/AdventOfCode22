package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int64
}
type Position struct {
	X, Y      int64
	Direction string
	Amount    int
}

/*
func MoveP1(x, y int64) {

	h.x = h.x + x
	h.y = h.y + y

	if h.x-t.x > 1 {
		t.x++
		t.y = h.y
	}
	if h.x-t.x < -1 {
		t.x--
		t.y = h.y

	}
	if h.y-t.y > 1 {
		t.y++
		t.x = h.x
	}
	if h.y-t.y < -1 {
		t.y--
		t.x = h.x

	}

	AddTailToTailspots(t.x, t.y)
}
*/
// Modified Move for P2 = Use an array, and limit moves to 1 row/col at a time
func Move(x, y int64) {

	Knots[0].x = Knots[0].x + x
	Knots[0].y = Knots[0].y + y
	for i := 1; i < len(Knots); i++ {

		if Knots[i-1].x-Knots[i].x > 1 {
			Knots[i].x++
			Knots[i].y = Knots[i].y + MaxMove(Knots[i-1].y-Knots[i].y)

		}
		if Knots[i-1].x-Knots[i].x < -1 {
			Knots[i].x--
			Knots[i].y = Knots[i].y + MaxMove(Knots[i-1].y-Knots[i].y)
		}
		if Knots[i-1].y-Knots[i].y > 1 {
			Knots[i].y++
			Knots[i].x = Knots[i].x + MaxMove(Knots[i-1].x-Knots[i].x)
		}
		if Knots[i-1].y-Knots[i].y < -1 {
			Knots[i].y--
			Knots[i].x = Knots[i].x + MaxMove(Knots[i-1].x-Knots[i].x)
		}

	}

	AddTailToTailspots(Knots[len(Knots)-1].x, Knots[len(Knots)-1].y)
}
func MaxMove(in int64) int64 {
	if in > 0 {
		return 1
	}
	if in < 0 {
		return -1
	}
	return 0
}

func AddTailToTailspots(x, y int64) {
	key := strconv.FormatInt(x, 10) + "_" + strconv.FormatInt(y, 10)
	tailspots[key] = 1

}

var tailspots map[string]int
var fname string
var Knots [10]Point

func main() {

	fname = "input.txt"
	//fname = "exampleinput.txt"

	tailspots = make(map[string]int)

	readFile, err := os.Open(fname)
	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		moveval, _ := strconv.Atoi(string(line[2:]))
		dir := string(line[0])
		if dir == "R" {

			for i := 0; i < moveval; i++ {
				Move(1, 0)
			}
		}
		if dir == "L" {
			for i := 0; i < moveval; i++ {
				Move(-1, 0)
			}
		}
		if dir == "U" {
			for i := 0; i < moveval; i++ {
				Move(0, 1)
			}
		}
		if dir == "D" {
			for i := 0; i < moveval; i++ {
				Move(0, -1)
			}
		}

	}

	fmt.Println(len(tailspots))
}
