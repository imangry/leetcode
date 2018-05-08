package algorithm_book

import (
	"testing"
	"fmt"
)

func TestLRUCache_Get(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Set(1, 1)
	cache.Set(2, 2)
	fmt.Println(cache.M)
	cache.Set(3,3)
	fmt.Println(cache.M)
}

