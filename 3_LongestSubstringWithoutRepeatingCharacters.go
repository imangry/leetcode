package main



func lengthOfLongestSubstring(s string) (max int) {
	var i, j int
	set := make(map[byte]interface{})

	for j < len(s) {
		if _, ok := set[s[j]]; !ok {
			set[s[j]] = struct{}{}
			j++
			if len(set) > max {
				max = len(set)
			}
		} else {
			delete(set, s[i])
			i++
		}
	}
	return max
}

func lengthOfLongestSubstring_1(s string) (max int) {
	m := make(map[byte]int)

	for i, j := 0, 0; j < len(s); j++ {
		if index, ok := m[s[j]]; ok {
			if i < index {
				i = index
			}
		}
		if j-i+1 > max {
			max = j - i + 1
		}
		m[s[j]] = j + 1
	}
	return max
}
