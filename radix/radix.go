package radix

import (
	"fmt"
	"math"
	"strings"
)

const (
	ALPHA_NUMBER = 27
)

var EPSILONE = "\000"

type Node struct {
	data     string
	leaf     bool
	children [ALPHA_NUMBER]*Node
}

type Tree struct {
	root *Node
}

func pprint(arg interface{}) {
	fmt.Println(arg)
}

func newNode(cont string) *Node {
	node := &Node{data: cont, leaf: true}
	for i := 0; i < ALPHA_NUMBER; i++ {
		node.children[i] = nil
	}
	return node
}

func (nd *Node) addSimpleNode(cont string) {
	nd.leaf = false

	node := &Node{data: cont, leaf: true}

	for i := 0; i < ALPHA_NUMBER; i++ {
		node.children[i] = nil
	}

	index := cont[0] - 'a'

	if cont == EPSILONE {
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
	newNd.leaf = nd.leaf
	return newNd
}

func (nd *Node) addComplexeNode(cont string) {
	index := cont[0] - 'a'
	if nd.children[index] == nil {
		if nd.data != EPSILONE && nd.leaf {
			nd.addSimpleNode(EPSILONE)
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
			nd.children[index].leaf = false
			nd.children[index].addSimpleNode(cont[match:])
			// pprint(nd.children[index])
		}
	}
}

func TreeInit() *Tree {
	head := newNode(EPSILONE)
	return &Tree{head}
}

func (nd *Node) addNode(cont string) {
	if nd.leaf {
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
	if nd.leaf {
		pprint(level + "-> " + appendix + nd.data)
	} else {
		for i := 0; i < ALPHA_NUMBER; i++ {
			if nd.children[i] != nil {
				nd.children[i].print(appendix+nd.data, level+level)
			}
		}
	}
}

func (nd *Node) SimplePrint(appendix string) {
	if nd.leaf {
		pprint(appendix + nd.data)
	} else {
		for i := 0; i < ALPHA_NUMBER; i++ {
			if nd.children[i] != nil {
				nd.children[i].SimplePrint(appendix + nd.data)
			}
		}
	}

}

func (nd Node) search(target string) bool {
	if target == "" && (nd.leaf || nd.children[26] != nil) {
		return true
	} else if target == "" {
		return false
	}

	index := target[0] - 'a'
	if nd.children[index] == nil {
		return false
	}

	if len(nd.children[index].data) > len(target) {
		return false
	}

	return nd.children[index].search(target[len(nd.children[index].data):])
}

func (tree Tree) SearchTree(target string) bool {
	// var path *[]string
	exists := tree.root.search(target)
	return exists
}

func (nd Node) getLastNode(target string, path *[]string) *Node {

	if target == "" {
		return &nd
	}

	index := target[0] - 'a'

	if nd.children[index] == nil {
		return nil
	}

	matched := countMatch(target, nd.children[index].data)

	*path = append(*path, target[:matched])

	if matched != len(nd.children[index].data) {

		return nd.children[index]
	}

	return nd.children[index].getLastNode(target[matched:], path)
}

// func checkPrefix(main string, sub string) (bool, string) {
// 	if main == sub {
// 		return false, ""
// 	}

// 	remainder := strings.Replace(main, sub, "", 1)

// 	return true, remainder
// }

func (nd Node) getPossibleSuffixes(start string) {
	path := []string{}
	branchStart := nd.getLastNode(start, &path)

	// prefixed, _ := checkPrefix(branchStart.data, path[len(path)-1])

	fullPath := strings.Join(path[:len(path)-1], "")
	// branchStart.print(fullPath, "--")

	branchStart.SimplePrint(fullPath)
}

func (tree Tree) AutoComplete(start string) bool {
	tree.root.getPossibleSuffixes(start)
	return true
}

func (tree Tree) Print() {
	tree.root.print("", "--")
}

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
	// tree.addword("abd")
	// pprint("---------- head ----------  ")
	// pprint(tree.root)
	// pprint("\n")
	// pprint(tree.root.children[10])
	// pprint("\n")
	// pprint(tree.root.children[10].children[0])
	// pprint("\n")
	// pprint(tree.root.children[2].children[0])
	// pprint(tree.root.children[3])

	// pprint(tree.SearchTree("czbaa"))

	tree.AutoComplete("c")
}
