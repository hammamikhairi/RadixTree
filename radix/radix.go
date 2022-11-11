package radix

import (
	"fmt"
	"math"
	// "strings"
)

const (
	ALPHA_NUMBER = 27
)

var EPS = "\000"

type Node struct {
	data     string
	end      bool
	children [ALPHA_NUMBER]*Node
}

type Tree struct {
	root *Node
}

func pprint(arg interface{}) {
	fmt.Println(arg)
}

func newNode(cont string) *Node {
	node := &Node{data: cont, end: true}
	for i := 0; i < ALPHA_NUMBER; i++ {
		node.children[i] = nil
	}
	return node
}

func (nd *Node) addSimpleNode(cont string) {
	nd.end = false
	node := &Node{data: cont, end: true}
	for i := 0; i < ALPHA_NUMBER; i++ {
		node.children[i] = nil
	}
	index := cont[0] - 'a'
	if cont == EPS {
		index = 26
	}
	if nd.children[index] != nil {
		panic("cant rewrite a node")
	}
	nd.children[index] = node
}

func countMatch(main string, sub string) int {
	min_length := math.Min(float64(len(main)), float64(len(sub)))
	counter := 0
	for i := 0; i < int(min_length); i++ {
		if main[i] == sub[i] {
			counter++
		} else {
			break
		}
	}
	return counter
}

func (nd *Node) cutNode(newData string) *Node {
	newNd := newNode(newData)
	flag := true
	for i := 0; i < ALPHA_NUMBER; i++ {
		if flag && nd.children[i] != nil {
			flag = false
		}
		newNd.children[i] = nd.children[i]
		nd.children[i] = nil
	}
	newNd.end = nd.end
	return newNd
}

func (nd *Node) addComplexeNode(cont string) {
	index := cont[0] - 'a'
	if nd.children[index] == nil {
		if nd.data != EPS && nd.end {
			nd.addSimpleNode(EPS)
		}
		nd.addSimpleNode(cont)
	} else {
		nodeData := nd.children[index].data
		match := countMatch(nodeData, cont)

		if match == len(nodeData) {
			nd.children[index].addComplexeNode(cont[match:])
		} else {
			newOrigin := nodeData[:match]
			newSuffix := nodeData[match:]
			newNd := nd.children[index].cutNode(newSuffix)
			nd.children[index].data = newOrigin
			nd.children[index].children[newSuffix[0]-'a'] = newNd
			nd.children[index].end = false
			nd.children[index].addSimpleNode(cont[match:])
			// pprint(nd.children[index])
		}
	}
}

func TreeInit() *Tree {
	head := newNode(EPS)
	return &Tree{head}
}

func (nd *Node) addNode(cont string) {
	if nd.end {
		nd.addSimpleNode(cont)
	} else {
		nd.addComplexeNode(cont)
	}
}

func (tree *Tree) Addword(word string) {
	tree.root.addNode(word)
}

func (nd *Node) print(appendix string, level string) {
	pprint(level + nd.data)
	if nd.end {
		pprint(level + "# " + appendix + nd.data)
	} else {
		for i := 0; i < ALPHA_NUMBER; i++ {
			if nd.children[i] != nil {
				nd.children[i].print(appendix+nd.data, level+"--")
			}
		}
	}

}

func (tree Tree) Print() {
	tree.root.print("", "--")
}

// func main() {
// 	// node := Node{data: "dd", end: false}

// 	// node.addSimpleNode("ddee")

// 	// pprint(node.children)

// 	tree := TreeInit()
// 	tree.Addword("aban")
// 	tree.Addword("cabi")
// 	tree.Addword("caba")
// 	tree.Addword("cabaa")
// 	tree.Addword("czbaa")
// 	tree.Addword("khairi")
// 	tree.Addword("khairis")
// 	tree.Addword("khkkris")
// 	tree.Addword("khkklis")
// 	// tree.addword("abd")
// 	// pprint("---------- head ----------  ")
// 	// pprint(tree.root)
// 	// pprint('\n')
// 	// pprint(tree.root.children[2])
// 	// pprint('\n')
// 	// pprint(tree.root.children[2].children[25])
// 	// pprint('\n')
// 	// pprint(tree.root.children[2].children[0])
// 	// pprint(tree.root.children[3])

// 	tree.Print()
// }
