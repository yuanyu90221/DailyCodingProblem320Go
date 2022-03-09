package find_max_distinct

func FindMaxDistinct(input string) int {
	charMap := make(map[rune]int)
	// maxStart := 0
	max := 0
	// curStart := 0
	var preChar rune
	curAccum := 0
	for idx, ch := range input {
		if _, ok := charMap[ch]; ok { // find existed char
			// curStart = val + 1
			if curAccum > max {
				max = curAccum
				// maxStart = curStart
			}
			if preChar == ch {
				curAccum = 1
				// curStart = idx
				charMap = make(map[rune]int)
			}
		} else {
			curAccum++
		}
		charMap[ch] = idx
		preChar = ch
	}
	if curAccum > max {
		max = curAccum
		// maxStart = curStart
	}
	return max
}
