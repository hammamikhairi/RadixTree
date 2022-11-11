package main

import (
	. "example/mymodule/Radixtree/radix"
	"fmt"
)

func main() {
	tree := TreeInit()

	// add words
	tree.Addword("aban")
	tree.Addword("cabi")
	tree.Addword("caba")
	tree.Addword("cabaa")
	tree.Addword("czbaa")
	tree.Addword("khairi")
	tree.Addword("khairis")
	tree.Addword("khkkris")
	tree.Addword("khkklis")

	// prints the tree (currently unhuman)
	tree.Print()

	// checks if a word exists
	fmt.Println(tree.SearchTree("k"))       // -> false
	fmt.Println(tree.SearchTree("kh"))      // -> false
	fmt.Println(tree.SearchTree("kha"))     // -> false
	fmt.Println(tree.SearchTree("khai"))    // -> false
	fmt.Println(tree.SearchTree("khair"))   // -> false
	fmt.Println(tree.SearchTree("khairi"))  // -> true
	fmt.Println(tree.SearchTree("khairis")) // -> true
}
