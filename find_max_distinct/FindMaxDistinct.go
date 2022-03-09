package find_max_distinct

func FindMaxDistinct(input string) int {
	charMap := make(map[rune]int)
	// maxStart := 0
	max := 0
	// curStart := 0
	curAccum := 0
	for idx, ch := range input {
		if _, ok := charMap[ch]; ok { // find existed char
			// curStart = val + 1
			if curAccum > max {
				max = curAccum
				// maxStart = curStart
			}
			// fmt.Println(curStart, maxStart)
		} else {
			curAccum++
		}
		charMap[ch] = idx
	}
	return max
}
