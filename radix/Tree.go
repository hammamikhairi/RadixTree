package radix

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Tree struct {
	Root *Node
}

// Initilizes a Radix Tree
func TreeInit() *Tree {
	head := newNode(EPSILONE)
	return &Tree{head}
}

// Adds a single word to the tree
//
// Parameters:
//   - `word` : the word to add
func (tree *Tree) Addword(word string) {
	tree.Root.addNode(word)
}

// Fills the tree with words from a csv file.
//
// Parameters:
//   - `path` : path to csv file
//
// **NOTE** each word must be in the first column of every row
func (tree *Tree) FillTree(path string) error {

	if !strings.HasSuffix(path, ".csv") {
		panic("can only read csv files (for now atleast)")
	}

	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()
	if err != nil {
		return nil
	}

	for _, row := range data {
		err = assertData(row[0])

		if err != nil {
			panic(fmt.Errorf("%w \nword : %s\n%s^", err, row[0], strings.Repeat(" ", digitPrefix(row[0])+7)))
		}

		tree.Addword(row[0])
	}

	return nil
}

func (tree Tree) SearchTree(target string) bool {
	// var path *[]string
	exists := tree.Root.search(target)
	return exists
}

func (tree Tree) AutoComplete(start string, PrintRes bool) int {

	if start == "" {
		pprint("Cant search empty strings")
		return 0
	}

	if start[0]-'a' > ALPHA_NUMBER-1 {
		return 0
	}

	printRes = PrintRes
	return tree.Root.getPossibleSuffixes(start)
}

func (tree Tree) Print() {
	tree.Root.print("", "--")
}

func (tree *Tree) RemoveWord(target string) error {

	// if target[0]-'a' > ALPHA_NUMBER-1 {
	// 	return
	// }

	// if !tree.SearchTree(target) {
	// 	pprint("doesnt belong")
	// 	return
	// }

	// tree.Root.removeNode(target)
	return errNotImplemented
}
