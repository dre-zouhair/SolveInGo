package cmd

func MinWindow(s string, t string) string {
	need := make(map[uint8]int)
	targetLen := 0
	for idx := range t {
		if _, ok := need[t[idx]]; !ok {
			targetLen++
		}
		need[t[idx]] += 1
	}

	window := make(map[uint8]int)

	left, right, valid := 0, 0, 0

	targetString := ""

	for right < len(s) {
		window[s[right]]++
		if need[s[right]] != 0 && window[s[right]] == need[s[right]] {
			valid++
		}
		right++
		for valid == targetLen {
			if targetString == "" {
				targetString = s[left:right]
			}
			if len(s[left:right]) < len(targetString) {
				targetString = s[left:right]
			}
			window[s[left]]--
			if need[s[left]] != 0 && window[s[left]] < need[s[left]] {
				valid--
			}
			left++
		}
	}
	return targetString
}
