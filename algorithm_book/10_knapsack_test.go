package algorithm_book

import (
	"testing"
	"fmt"
)

func TestKnapsackWithoutRepetition(t *testing.T) {
	maxWeight := 10

	weight := []int{3, 4, 5}
	value := []int{4, 6, 7}
	fmt.Println(KnapsackWithoutRepetition(maxWeight, weight, value))
	fmt.Println(KnapsackWithoutRepetition_1(maxWeight, weight, value))
}
