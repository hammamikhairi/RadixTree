package main

import (
	"bufio"
	"fmt"
	Radix "github.com/hammamikhairi/RadixTree/radix"
	"os"
	"strings"
	"time"
)

// This is an Example use of the package
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
