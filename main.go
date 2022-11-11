package main

import (
	. "example/mymodule/Radixtree/radix"
)

func main() {
	tree := TreeInit()

	tree.Addword("aban")
	tree.Addword("cabi")
	tree.Addword("caba")
	tree.Addword("cabaa")
	tree.Addword("czbaa")
	tree.Addword("khairi")
	tree.Addword("khairis")
	tree.Addword("khkkris")
	tree.Addword("khkklis")

	tree.Print()
}
