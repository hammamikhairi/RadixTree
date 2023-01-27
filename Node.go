package RadixTree

import "strings"

type Node struct {
	data     string
	leaf     bool
	Children [ALPHA_NUMBER]*Node
}

func (nd *Node) addSimpleNode(cont string) {
	nd.leaf = false

	node := &Node{data: cont, leaf: true}

	for i := 0; i < ALPHA_NUMBER; i++ {
		node.Children[i] = nil
	}

	var index int = getIndex(cont)

	if nd.Children[index] != nil {
		return
	}

	nd.Children[index] = node
}

func (nd *Node) cutNode(newData string) *Node {
	newNd := newNode(newData)
	flag := true
	for i := 0; i < ALPHA_NUMBER; i++ {
		if flag && nd.Children[i] != nil {
			flag = false
		}
		newNd.Children[i] = nd.Children[i]
		nd.Children[i] = nil
	}
	newNd.leaf = nd.leaf
	return newNd
}

func (nd *Node) addComplexeNode(cont string) {
	var index int = getIndex(cont)

	if nd.Children[index] == nil {
		if nd.data != EPSILONE && nd.leaf {
			nd.addSimpleNode(EPSILONE)
		}
		nd.addSimpleNode(cont)

	} else if cont == nd.Children[index].data {
		nd.Children[index].addSimpleNode(EPSILONE)
	} else {
		nodeData := nd.Children[index].data
		match := countMatch(nodeData, cont)

		if match == len(nodeData) {
			nd.Children[index].addComplexeNode(cont[match:])
		} else {
			newOrigin := nodeData[:match]
			newSuffix := nodeData[match:]
			newNd := nd.Children[index].cutNode(newSuffix)
			nd.Children[index].data = newOrigin
			nd.Children[index].Children[getIndex(newSuffix)] = newNd
			nd.Children[index].leaf = false
			nd.Children[index].addSimpleNode(cont[match:])
		}
	}
}

func (nd *Node) addNode(cont string) {
	if nd.leaf {
		nd.addSimpleNode(cont)
	} else {
		nd.addComplexeNode(cont)
	}
}

func (nd *Node) print(appendix string, level string) {
	pprint(level + nd.data)
	if nd.leaf {
		pprint(level + "-> " + appendix + nd.data)
	} else {
		for i := 0; i < ALPHA_NUMBER; i++ {
			if nd.Children[i] != nil {
				nd.Children[i].print(appendix+nd.data, level+"--")
			}
		}
	}
}

func (nd *Node) SimplePrint(appendix string, wordsCount *int) {
	if nd.leaf {
		*wordsCount += 1
		if printRes {
			pprint(appendix + nd.data)
		}
	} else {
		for i := 0; i < ALPHA_NUMBER; i++ {
			if nd.Children[i] != nil {
				nd.Children[i].SimplePrint(appendix+nd.data, wordsCount)
			}
		}
	}

}

func (nd Node) search(target string) bool {
	if target == "" && (nd.leaf || nd.Children[26] != nil) {
		return true
	} else if target == "" {
		return false
	}

	var index int = getIndex(target)
	if nd.Children[index] == nil {
		return false
	}

	if len(nd.Children[index].data) > len(target) {
		return false
	}

	return nd.Children[index].search(target[len(nd.Children[index].data):])
}

func (nd Node) getLastNode(target string, path *[]string) *Node {

	if target == "" {
		return &nd
	}

	var index int = getIndex(target)

	if nd.Children[index] == nil {
		return nil
	}

	matched := countMatch(target, nd.Children[index].data)

	*path = append(*path, target[:matched])

	if matched == 0 {
		return &nd
	}

	return nd.Children[index].getLastNode(target[matched:], path)
}

func (nd Node) getPossibleSuffixes(start string) int {
	path := []string{}
	branchStart := nd.getLastNode(start, &path)

	var fullPath string
	if len(path) != 0 {
		fullPath = strings.Join(path[:len(path)-1], "")
	} else {
		fullPath = ""
	}

	if branchStart == nil {
		pprint("Word doesnt exist")
	} else {
		var wordsCount int = 0
		branchStart.SimplePrint(fullPath, &wordsCount)
		return wordsCount
	}
	return 0
}

// func (nd *Node) wipe() {
// 	nd.data = "aaaaa"
// 	if nd.leaf {
// 		nd = nil
// 	} else {
// 		nd.Children[ALPHA_NUMBER-1] = nil
// 	}
// }

// func (nd *Node) removeNode(target string) error {

// 	index := target[0] - 'a'
// 	matched := countMatch(target, nd.Children[index].data)
// 	newTarget := target[matched:]
// 	newMatched := countMatch(newTarget, nd.Children[index].data)

// 	pprint(matched)
// 	pprint(target)
// 	pprint(newTarget)
// 	pprint(newMatched)

// 	if newTarget[newMatched:] == "" {
// 		if nd.Children[index].leaf {
// 			// remove from nd
// 			hasChildren := false
// 			nd.Children[index] = nil

// 		childrenCheck:
// 			for i := 0; i < ALPHA_NUMBER; i++ {
// 				if nd.Children[i] != nil {
// 					hasChildren = true
// 					break childrenCheck
// 				}
// 			}

// 			if !hasChildren {
// 				nd.leaf = true
// 			}

// 			return nil

// 		} else {
// 			// remove eub
// 			nd.Children[index].Children[newTarget[0]-'a'].Children[ALPHA_NUMBER-1] = nil
// 			// pprint(nd)
// 			// pprint(nd.Children[index])
// 			childrenNum := 0
// 			var cursor int

// 			for i := 0; i < ALPHA_NUMBER; i++ {
// 				if nd.Children[index].Children[newTarget[0]-'a'].Children[i] != nil {
// 					childrenNum++
// 					cursor = i
// 				}
// 			}

// 			if childrenNum == 0 {
// 				nd.Children[index].Children[newTarget[0]-'a'].leaf = true
// 			} else if childrenNum == 1 {
// 				toAppend := nd.Children[index].Children[newTarget[0]-'a'].Children[cursor]
// 				nd.Children[index].data += toAppend.data
// 				// pprint(nd.Children[index].data)
// 				for i := 0; i < ALPHA_NUMBER; i++ {
// 					if cursor != i {
// 						nd.Children[index].Children[newTarget[0]-'a'].Children[i] = nd.Children[index].Children[newTarget[0]-'a'].Children[cursor].Children[i]
// 					}
// 				}

// 				nd.Children[index].Children[newTarget[0]-'a'].leaf = nd.Children[index].Children[newTarget[0]-'a'].Children[cursor].leaf
// 				nd.Children[index].Children[newTarget[0]-'a'].Children[cursor] = nil
// 				// nd.leaf = true
// 			}
// 			return nil
// 		}
// 	}

// 	nd.Children[index].removeNode(newTarget)
// 	return nil
// }
