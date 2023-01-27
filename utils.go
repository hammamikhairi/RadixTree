package radixtree

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"unicode"
)

const (
	ALPHA_NUMBER      = 48
	EPSILONE          = "\000"
	ACCEPTED_SYMBOLES = " \"'(),;:!?.1234567890"
)

var SymbolesToIndex map[string]int

var errNotImplemented = errors.New("not implemented yet")
var printRes = true

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

func fillSymbolesMap() map[string]int {
	alpha := 27
	STI := make(map[string]int)
	for index, symb := range ACCEPTED_SYMBOLES {
		STI[string(symb)] = alpha + index
	}

	for i := 'a'; i <= 'z'; i++ {
		STI[string(i)] = int(i) - 97
	}

	STI[EPSILONE] = 26

	return STI
}

func getIndex(letter string) (index int) {

	if letter == "" || letter == EPSILONE {
		index = 26
	} else {
		index = SymbolesToIndex[string(letter[0])]
	}
	return
}
