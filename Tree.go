package radixtree

import (
	"bufio"
	"os"
)

type Tree struct {
	Root *Node
}

// Initilizes a Radix Tree
func TreeInit() *Tree {
	SymbolesToIndex = fillSymbolesMap()
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
//   - `fileStream` : *os.File
func (tree *Tree) FillTree(fileStream *os.File) error {

	fileScanner := bufio.NewScanner(fileStream)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		tree.Addword(fileScanner.Text())
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
	defer func() {
		printRes = true
	}()

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
