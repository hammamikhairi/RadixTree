package RadixTree

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
)

func getData() [][]string {
	file, err := os.Open("../data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	return data
}

var Data [][]string = getData()

func TestWordsAddition(t *testing.T) {
	tree := TreeInit()
	data := getData()
	fmt.Println(ALPHA_NUMBER)
	// tree.Addword("bbs")
	// // add words
	for _, line := range data {
		tree.Addword(line[0])
	}

DataLoop:
	for _, line := range data {
		if !tree.SearchTree(line[0]) {
			t.Errorf("Word <%s> not found", line[0])
			break DataLoop
		}
	}

	// tree.Print()
}

// func TestEee(t *testing.T) {
// 	t.Error("two")
// }

// func TestKjb(t *testing.T) {
// 	t.Error("three")
// }
