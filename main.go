package main

import (
	Radix "RadixTree/radix"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	tree := Radix.TreeInit()
	fileStream, err := os.Open("SearchQueriesSample.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer fileStream.Close()

	tree.FillTree(fileStream)

	// Testing
	var input string
	var start time.Time
	var end time.Duration
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("~ search : ")

		input, err = in.ReadString('\n')
		if err != nil {
			panic(err)
		}

		start = time.Now()
		number := tree.AutoComplete(strings.Replace(input, "\n", "", 1), true)
		end = time.Since(start)
		fmt.Printf("%d matches found in ", number)
		fmt.Println(end)
		input = ""
	}

}
