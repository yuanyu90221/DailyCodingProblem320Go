# DailyCodingProblem320Go

## 題目描述

給定一個字串，找出最小長度的子字串使得內部所有的字元在子字串內都只出現過一次。

舉例來說，給定字串 "jiujitsu" ，則回傳 5 ， 因為 "jitsu" 內每個字元都只出現過一次。

## 解題思路


要找出包含最多不同字元的子字串。

初始化最大長度為 0

初始化子字串起始位置為 0

要找的子字串有兩個特性

1 字元位置是連續的

2 字元都是唯一出現

可以透過每次把出現過的字串紀錄在 Map ， Map 以該字元為 key 值是出現的位置

設定當下累計長度初始為 = 0，當下起始位置為 0

以每個字元當作字串起點每次檢查 Map 是否出現過

如果沒有出現，把累計長度加 1 

如果出現過，則比較最大長度與累計長度

需要注意如果連續出現兩個一樣的字元，需要把開始位址更新為目前的位置，並且設定累計長度為 1，並且清除原本累計 Map

如果最大長度小於累計長度，則把最大長度設為累計長度，更新字串起位置為新為該 Map 該字元對應到的值 +1 

起始位置更新為該 Map 該字元對應到的值 +1，並且更新遇到重複字元的位置為當下位置

當走訪完所有字元, 最大長度就是答案

## 程式碼

```golang
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
```
