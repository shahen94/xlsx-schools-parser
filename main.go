package main

import (
	"fmt"
	"os"
	"path"

	"actogram.net/go-xlsx-parser/computing"
	"actogram.net/go-xlsx-parser/plt4m"
	"github.com/tealeg/xlsx"
)

var states *computing.StateMapper

func log(msg string) {
	fmt.Println("[LOG] ", msg)
}

func openXLSX(filePath string) *plt4m.SchoolList {
	dir, _ := os.Getwd()
	file, err := xlsx.OpenFile(path.Join(dir, filePath))

	if err != nil {
		panic(err)
	}

	return &plt4m.SchoolList{file}
}

func init() {
	states = &computing.StateMapper{
		Path: "states.csv",
	}

	err := states.Read()

	if err != nil {
		panic(err)
	}
}

func main() {
	log("Loading file...")
	file := openXLSX("schools.xlsx")
	file.PrintCount()

	log("Adding new column...")
	file.AddColumn("Offset")

	log("Iterating over rows...")
	file.IterateOver(computing.ComputeOffset(states))

	log("Saving to the file...")
	file.WriteUpdates("edited.xlsx")

	log("Done!")
}
