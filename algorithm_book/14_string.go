package algorithm_book


/*
求两个字符串是不是变位词

解：
通过额外的数组保存词频
时间复杂度 O(n) 空间复杂度 O(1)
*/
func TwoStringsAreAnagrams(s1, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 || len(s1) != len(s2) {
		return false
	}

	letters := [256]int{}

	for i := 0; i < len(s1); i++ {
		letters[s1[i]] ++
		letters[s2[i]]--
	}

	for _, value := range letters {
		if value != 0 {
			return false
		}
	}
	return true
}
