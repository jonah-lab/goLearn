package main

func findNoDiffStr(s string) int {
	lastoccued := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastoccued[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 >= maxLength {
			maxLength = i - start + 1
		}
		lastoccued[ch] = i
	}
	return maxLength
}
