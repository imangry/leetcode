package basic_algorithm

import (
	"testing"
	"math/rand"
	"time"
	"os"
)

func TestNewRBNode(t *testing.T) {
	rand.Seed(time.Now().Unix())
	root := RBInit()
	for i := 0; i < 16; i++ {
		root = RBInsert(root, rand.Int()%50)
	}
	PrintRBTree(root,os.Stdout)
}
