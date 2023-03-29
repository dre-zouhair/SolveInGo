package cmd

import (
	"Blind75MustDo/structs"
	"strings"
)

func LongestCommonSubsequence(text1 string, text2 string) int {
	length1, length2 := len(text1), len(text2)
	dp := make([][]int, length1+1)
	for i := 0; i <= length1; i++ {
		dp[i] = make([]int, length2+1)
	}
	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i][j-1] > dp[i-1][j] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[length1][length2]
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make(map[structs.DP]int)
	return longestCommonSubsequenceRec(text1, text2, 0, dp)
}

func longestCommonSubsequenceRec(text1 string, text2 string, length int, dp map[structs.DP]int) int {

	el := structs.DP{Text1: text1, Text2: text2}

	if _, ok := dp[el]; ok {
		return dp[el]
	}

	if len(text1) == 0 || len(text2) == 0 {
		dp[el] = length
		return length
	}

	if len(text1) == 1 && len(text2) == 1 && text1 != text2 {
		dp[el] = length
		return length
	}

	if strings.Contains(text1, text2) {
		return length + len(text2)
	}

	if strings.Contains(text2, text1) {
		return length + len(text1)
	}

	a, b, c := 0, 0, 0
	t1, t2 := len(text1), len(text2)
	if text1[0] == text2[0] {
		a = longestCommonSubsequenceRec(text1[1:t1], text2[1:t2], length+1, dp)
	}

	b = longestCommonSubsequenceRec(text1[1:t1], text2, length, dp)
	c = longestCommonSubsequenceRec(text1, text2[1:t2], length, dp)

	if a > b && a > c && a > length {
		length = 1 + a
	}

	if b > a && b > c && b > length {
		length = b
	}

	if c > a && c > b && c > length {
		length = c
	}
	dp[el] = length
	return length

}
