package radix

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"unicode"
)

// TODO : full phrases support ( add spaces, numbers, commas, ... to the alphabet )
// TODO : write tests XD
// TODO : Think of better Printing
// TODO : add flags

const (
	ALPHA_NUMBER = 27
	EPSILONE     = "\000"
)

var errNotImplemented = errors.New("not implemented yet!")

type Node struct {
	data     string
	leaf     bool
	Children [ALPHA_NUMBER]*Node
	// parent   *Node
}

func pprint(arg interface{}) {
	fmt.Println(arg)
}

func newNode(cont string) *Node {
	node := &Node{data: cont, leaf: true}
	for i := 0; i < ALPHA_NUMBER; i++ {
		node.Children[i] = nil
	}
	return node
}

func (nd *Node) addSimpleNode(cont string) {
	nd.leaf = false

	node := &Node{data: cont, leaf: true}

	for i := 0; i < ALPHA_NUMBER; i++ {
		node.Children[i] = nil
	}

	var index int

	if cont == "" || cont == EPSILONE {
		index = 26
	} else {
		index = int(cont[0] - 'a')
	}

	if nd.Children[index] != nil {
		panic("cant rewrite a node")
	}

	nd.Children[index] = node
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

func assertData(input string) error {

	numeric := regexp.MustCompile(`^[A-Za-z]*$`).MatchString(input)

	if !numeric {
		return errors.New("unsupported word")
	}

	return nil
}

func digitPrefix(s string) int {
	for i, r := range s {
		if unicode.IsDigit(r) {
			return i
		}
	}
	return -1
}
