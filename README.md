# Radix Tree



## What is a Radix Tree

a radix tree (also radix trie or compact prefix tree) is a data structure that represents a space-optimized trie in which each node that is the only child is merged with its parent.

<p align="center">
  <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Patricia_trie.svg/400px-Patricia_trie.svg.png" alt="Radix Tree"/>
</p>

## Usage

```go

package main

import (
	Radix "RadixTree/radix"
)

func main() {

  // initialize tree
  tree := Radix.TreeInit()

  // add a word to the tree
  tree.Addword("khairi")

  // fill the tree with words from a csv file
  err := tree.FillTree("data.csv")

  if err != nil {
    panic(err)
  }

  // search for a word in the tree
  found := tree.SearchTree("khairi")

  // auto complete a given sub string
  propNumber := tree.AutoComplete(input, false)

}

```

## Project52

This project is a part of [LFSA 88](https://github.com/hammamikhairi/LFSA-88).


## License

This package is licensed under MIT license. See LICENSE for details.
