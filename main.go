package main

import (
	. "example/mymodule/Radixtree/radix"
	"fmt"
)

func main() {

	//initialize tree
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
	fmt.Println("------- Print Tree -------")
	tree.Print()

	// checks if a word exists
	fmt.Println("------- Word Checks -------")
	fmt.Println(tree.SearchTree("k"))       // -> false
	fmt.Println(tree.SearchTree("kh"))      // -> false
	fmt.Println(tree.SearchTree("kha"))     // -> false
	fmt.Println(tree.SearchTree("khai"))    // -> false
	fmt.Println(tree.SearchTree("khair"))   // -> false
	fmt.Println(tree.SearchTree("khairi"))  // -> true
	fmt.Println(tree.SearchTree("khairis")) // -> true

	// auto completion based on initial strings (not finalized ( no checks ) )
	fmt.Println("------- Auto Completion \"c\" -------")
	tree.AutoComplete("c")

	fmt.Println("------- Auto Completion \"kh\" -------")
	tree.AutoComplete("kh")
}
