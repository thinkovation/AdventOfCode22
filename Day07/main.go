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
func (fs FS) sizeofcontents(path string) int {
	var tot int
	for _, v := range fs.Entries {
		if len(v.Path) >= len(path) {
			//fmt.Println(v.Path, path)
			if string(v.Path[:len(path)]) == path {
				//	fmt.Println("Matched ", v.Path, " with ", path)
				tot = tot + v.Size

			}
		}
	}
	return tot
}

var Root FS

func main() {
	fname := "input.txt"
	//fname := "exampleinput.txt"
	Root.CurrentPath = "/"
	newVal := Object{
		Name:        "",
		Size:        0,
		Path:        Root.CurrentPath,
		IsDirectory: true,
	}

	Root.Entries = append(Root.Entries, newVal)
	readFile, err := os.Open(fname) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
		if len(line) > 0 {
			if string(line[0]) == "$" {
				fmt.Println("Command")
				if string(line[2:4]) == "cd" {
					fmt.Println("Change Directory")
					Root.CD(line[5:])
					fmt.Println(Root.CurrentPath)
				}
			} else {
				fmt.Println("File")
				var newVal Object
				els := strings.Split(line, " ")
				if string(line[0]) == "d" {
					newVal = Object{
						Name:        els[1],
						Size:        0,
						Path:        Root.CurrentPath,
						IsDirectory: true,
					}
				} else {

					fs, _ := strconv.Atoi(els[0])
					newVal = Object{
						Name:        els[1],
						Size:        fs,
						Path:        Root.CurrentPath,
						IsDirectory: false,
					}
				}
				Root.Entries = append(Root.Entries, newVal)

			}
		}

	}
	totfs := 0
	for _, v := range Root.Entries {
		//fmt.Println( v.Name, v.Size, v.Path)
		fmt.Println(v.Path+"-"+v.Name, " DIR?", v.IsDirectory)
		totfs = totfs + v.Size
	}
	fmt.Println("Total Files", totfs)
	var selfs int
	for _, v := range Root.Entries {
		if v.IsDirectory {
			fmt.Println(v.Path+v.Name, " DIR?", v.IsDirectory)
			if v.Name != "" {
				v.Name = v.Name + "/"
			}
			fs := Root.sizeofcontents(v.Path + v.Name)
			fmt.Println(fs)
			if fs <= 100000 {
				selfs = selfs + fs
			}

		}
		//fmt.Println( v.Name, v.Size, v.Path)

	}
	fmt.Println(selfs)

	// Part 2
	// We'll repeat the scan above but pop the results into a map of the path and the size
	const TotalCapacity = 70000000
	const CapacityRequired = 30000000
	var CapacityAvailable = TotalCapacity - totfs
	fmt.Println("Available Capacity", CapacityAvailable)
	var SpaceNeeded = CapacityRequired - CapacityAvailable
	fmt.Println(SpaceNeeded)
	var CandidateDirName string
	var CandidateDirSize int = totfs
	for _, v := range Root.Entries {
		if v.IsDirectory {
			fmt.Println(v.Path+v.Name, " DIR?", v.IsDirectory)
			if v.Name != "" {
				v.Name = v.Name + "/"
			}
			fs := Root.sizeofcontents(v.Path + v.Name)
			if fs > SpaceNeeded && fs < CandidateDirSize {
				CandidateDirSize = fs
				CandidateDirName = v.Path + v.Name
			}

		}
		//fmt.Println( v.Name, v.Size, v.Path)

	}
	fmt.Println("Dirname", CandidateDirName)
	fmt.Println("Dir Size", CandidateDirSize)

}
