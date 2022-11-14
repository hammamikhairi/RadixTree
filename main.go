package main

import (
	"encoding/csv"
	Radix "example/mymodule/Radixtree/radix"
	"fmt"
	"os"
	"time"
)

func getData() [][]string {
	f, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return data
}

func main() {

	data := getData()

	//initialize tree
	tree := Radix.TreeInit()

	for i, line := range data {
		if i > 0 {
			tree.Addword(line[0])
		}
	}

	// prints the tree (currently unhuman)
	// fmt.Println("------- Print Tree -------")
	// tree.Print()

	// auto completion based on initial strings

	fmt.Println("------- Confirm that every word in the dataset is in the tree -------")
	for i, line := range data {
		if i > 0 {
			if !tree.SearchTree(line[0]) {
				panic("A word is missing")
			}
		}
	}
	fmt.Println("All checks are good")

	fmt.Println("------- Auto Completion (no pre checks implemented yet) -------")
	var input string
	var start time.Time
	var end time.Duration
	for {
		fmt.Print("~ search : ")
		fmt.Scanf("%s", &input)

		// fmt.Print("------- Auto Completing \"")
		// fmt.Print(input)
		// fmt.Print("\" -------\n")

		start = time.Now()
		number := tree.AutoComplete(input)
		end = time.Since(start)
		fmt.Printf("%d matches found in ", number)
		fmt.Println(end)
	}
}
