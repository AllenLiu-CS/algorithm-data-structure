package rbtree

import (
	"log"
	"testing"
)

func TestRbtree(t *testing.T) {
	rbtree := NewRBTree()

	datas := []int64{1, 2, 3, 4, 5, 6, 7, 8}

	for _, num := range datas {
		rbtree.Insert(num)
	}

	log.Print("**********输出红黑树**********")
	printTreeInLog(rbtree.root, "(root)")
}
