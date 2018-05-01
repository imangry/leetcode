package algorithm_book

import (
	"math/rand"
	"time"
)
/*
洗牌算法
这个有点像概率中的 A几几
*/
func Shuffle(list []int) []int {
	if len(list) < 2 {
		return list
	}
	rand.Seed(time.Now().Unix())

	for i := len(list) - 1; i > 0; i-- {
		//因为要包括i的值，所以得加1，取余的时候才能余i
		u := rand.Uint64() % uint64(i+1)
		list[i], list[u] = list[u], list[i] //将随机得到的索引和最后一个交换
	}
	return list
}
