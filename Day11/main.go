package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items          []int
	Operator       string
	Operand        int
	TestDivisor    int
	TestFalse      int
	TestTrue       int
	ItemsInspected int
}

func (m *Monkey) AddItem(value int) {
	m.Items = append(m.Items, value)
}

func LoadMonkey(lines []string) {
	var newMonkey Monkey
	if len(lines) == 6 {
		//Skip line 1 - they're ordered from 0 so we can just append monkeys
		// StartingItems
		StartingItems := strings.Split(lines[1], " ")

		for i := 4; i < len(StartingItems); i++ {
			itemtoadd, err := strconv.Atoi(strings.Replace(StartingItems[i], ",", "", -1))
			if err != nil {
				fmt.Println(err.Error())
			} else {
				newMonkey.Items = append(newMonkey.Items, itemtoadd)

			}

		}

		//Operator and Operand
		OperationElements := strings.Split(lines[2], " ")

		newMonkey.Operator = OperationElements[6]
		if OperationElements[7] == "old" {
			newMonkey.Operand = 0
		} else {
			operand, _ := strconv.Atoi(OperationElements[7])
			newMonkey.Operand = operand

		}
		// Test
		testline := strings.Split(lines[3], " ")

		divisor, _ := strconv.Atoi(testline[5])
		newMonkey.TestDivisor = divisor
		// Test True
		testtrue := strings.Split(lines[4], " ")

		truethrow, _ := strconv.Atoi(testtrue[9])
		newMonkey.TestTrue = truethrow
		// Test False
		testfalse := strings.Split(lines[5], " ")

		falsethrow, _ := strconv.Atoi(testfalse[9])
		newMonkey.TestFalse = falsethrow

	}

	Monkeys = append(Monkeys, newMonkey)
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func lowestCommonMultiple(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = lowestCommonMultiple(result, integers[i])
	}

	return result
}

var LCM int

func MonkeyBusinessRestrained(id int) {
	for _, v := range Monkeys[id].Items {
		Monkeys[id].ItemsInspected++
		var ItemValue int
		if Monkeys[id].Operator == "+" {
			if Monkeys[id].Operand == 0 {
				ItemValue = v + v

			} else {
				ItemValue = v + Monkeys[id].Operand
			}
		} else {
			if Monkeys[id].Operand == 0 {
				ItemValue = v * v

			} else {
				ItemValue = v * Monkeys[id].Operand
			}
		}
		ItemValue = ItemValue / WorryDivisor
		if ItemValue%Monkeys[id].TestDivisor == 0 {
			Monkeys[Monkeys[id].TestTrue].AddItem(ItemValue)
		} else {
			Monkeys[Monkeys[id].TestFalse].AddItem(ItemValue)
		}
	}
	Monkeys[id].Items = []int{}

}
func MonkeyBusiness(id int) {
	for _, v := range Monkeys[id].Items {
		Monkeys[id].ItemsInspected++
		var ItemValue int
		if Monkeys[id].Operator == "+" {
			if Monkeys[id].Operand == 0 {
				ItemValue = v + v

			} else {
				ItemValue = v + Monkeys[id].Operand
			}
		} else {
			if Monkeys[id].Operand == 0 {
				ItemValue = v * v

			} else {
				ItemValue = v * Monkeys[id].Operand
			}
		}

		ItemValue = ItemValue / WorryDivisor
		if WorryDivisor == 1 {
			ItemValue = ItemValue % LCM
		}

		if ItemValue%Monkeys[id].TestDivisor == 0 {
			Monkeys[Monkeys[id].TestTrue].AddItem(ItemValue)
		} else {
			Monkeys[Monkeys[id].TestFalse].AddItem(ItemValue)
		}
	}
	Monkeys[id].Items = []int{}

}

var Monkeys []Monkey

//var WorryDivisor = 3
var WorryDivisor = 1

//var Iterations = 20
var Iterations = 10000

func main() {
	fname := "input.txt"
	//fname := "exampleinput.txt"

	readFile, err := os.Open(fname) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())

	}

	for i := 0; i < len(lines); i = i + 7 {

		LoadMonkey(lines[i : i+6])

	}
	// Find the Lowest common multiple
	var mults []int
	for i := 2; i < len(Monkeys); i++ {
		mults = append(mults, Monkeys[i].TestDivisor)
	}
	LCM = lowestCommonMultiple(Monkeys[0].TestDivisor, Monkeys[1].TestDivisor, mults...)
	fmt.Println(LCM)
	for i := 1; i <= Iterations; i++ {

		for k := range Monkeys {

			MonkeyBusiness(k)
		}
	}
	for k, v := range Monkeys {
		blah, _ := json.MarshalIndent(v, " ", " ")
		fmt.Println("Monkey", k, "-", string(blah))
	}
	var ItemsInspectedArray []int
	for k, v := range Monkeys {
		fmt.Println("K", k, "Times", v.ItemsInspected)
		ItemsInspectedArray = append(ItemsInspectedArray, v.ItemsInspected)

	}
	sort.Ints(ItemsInspectedArray[:])
	fmt.Println(ItemsInspectedArray)
	fmt.Println(ItemsInspectedArray[len(ItemsInspectedArray)-1] * ItemsInspectedArray[len(ItemsInspectedArray)-2])
	fmt.Println("Done ")

}
