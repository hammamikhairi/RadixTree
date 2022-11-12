package radix

import (
	"testing"
)

func TestPrintTree(t *testing.T) {
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

}

func TestEee(t *testing.T) {
	t.Error("two")
}

func TestKjb(t *testing.T) {
	t.Error("three")
}
