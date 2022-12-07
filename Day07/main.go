package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Object struct {
	Name        string
	Size        int
	Path        string
	IsDirectory bool
}
type FS struct {
	Entries     []Object
	CurrentPath string
}

func (fs *FS) CD(path string) {
	fmt.Println("PATH :", "@"+path+"@")
	if path == ".." {
		var newcp string
		sp := strings.Split(fs.CurrentPath, "/")
		fmt.Println("Cirrent path components", sp)
		for i := 0; i < len(sp)-2; i++ {
			fmt.Println(sp[i])
			newcp = newcp + sp[i] + "/"

		}
		fs.CurrentPath = newcp
		return
	}
	if string(path[0]) == "/" {
		fmt.Println("AHAA")
		fs.CurrentPath = path
		return
	}
	fmt.Println("WOOP")
	fs.CurrentPath = fs.CurrentPath + path + "/"
}

var Root FS

func main() {
	//fname := "input.txt"
	fname := "exampleinput.txt"
	Root.CurrentPath = "/"
	readFile, err := os.Open(fname) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)

		if string(line[:4]) == "$ ls" {
			fmt.Println("List Directory")
			fileScanner.Scan()
			line = fileScanner.Text()
			if line != "" {
				fmt.Println("LIne = ", "#"+line+"#")
				for string(line[0]) != "$" {
					fmt.Println(line)
					fileScanner.Scan()
					line = fileScanner.Text()
					fmt.Println("LIne = ", "#"+line+"#")
					fmt.Println("LoL:", len(line))
					if len(line) == 0 {
						line = "$"
						break
					}
					if string(line[0]) != "$" && string(line[:2]) != "dir" {
						els := strings.Split(line, " ")
						fmt.Println("Length:", els[0], "Name ", els[1], "Fullpath", Root.CurrentPath+els[1])
						intVar, _ := strconv.Atoi(els[0])
						newitem := Object{
							Name:        els[1],
							Size:        intVar,
							Path:        Root.CurrentPath + els[1],
							IsDirectory: false,
						}
						Root.Entries = append(Root.Entries, newitem)
					}

				}
			}

		}
		if line != "$" && string(line[:4]) == "$ cd" {
			fmt.Println("Change Directory to", string(line[5:]))
			Root.CD(string(line[5:]))
			fmt.Println("Current Directory :", Root.CurrentPath)
		}

	}
	for _, v := range Root.Entries {
		//fmt.Println( v.Name, v.Size, v.Path)
		fmt.Println(v.Path)
	}
}
