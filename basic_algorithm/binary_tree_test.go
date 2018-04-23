package basic_algorithm

import (
	"testing"
	"math/rand"
	"time"
	"os"
)

func TestNewTreeNode(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var root *TreeNode
	for i := 0; i < 10; i++ {
		root = InsertTreeNode(root, rand.Int()%20)
	}
	PrintTree(root,os.Stdout)
}
