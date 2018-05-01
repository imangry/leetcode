package dp

import (
	"testing"
	"fmt"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	list := []int{5, 3, 4, 8, 6, 7, 2}

	for i := 1; i <= len(list); i++ {
		fmt.Println(LongestIncreasingSubsequence(list[:i]))
	}

}
