package binaryindextree

import (
	"testing"
)

func TestInit(t *testing.T) {
	bits := &BinaryIndexedTree{}
	var initarr []int
	for i := 1; i <= 20; i++ {
		initarr = append(initarr, i)
	}
	bits.Init(initarr)
	t.Logf("tree is %v", bits.tree)
	t.Logf("Query(%d) is %d", 20, bits.Query(20))
	bits.Add(8, 1)
	t.Logf("tree is %v", bits.tree)
}
