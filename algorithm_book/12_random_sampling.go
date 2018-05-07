package algorithm_book

import "math/rand"

/*
从n大小的数组中随机取出m个值
有点像 C几几

洗牌算法是排列
抽样算法是组合

*/

func RandomSample(list []int, m int) []int {

	if len(list) < 1 || m <= 0 {
		return []int{}
	}

	result := make([]int, m)

	copy(result, list[:m])

	for i := 0; i < len(list); i++ {
		k := rand.Int63n(int64(i + 1))
		if k < int64(m) {
			result[k] = list[i]
		}

	}
	return result
}
