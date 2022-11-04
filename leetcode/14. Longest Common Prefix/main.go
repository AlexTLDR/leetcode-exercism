package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	lenStrings := len(strs)
	if lenStrings == 0 {
		return ""
	}

	firstString := strs[0]
	lenFirstString := len(firstString)

	commonPrefix := ""
	for i := 0; i < lenFirstString; i++ {
		firstStringRune := string(firstString[i])
		match := true
		for j := 1; j < lenStrings; j++ {
			if (len(strs[j]) - 1) < i {
				match = false
				break
			}

			if string(strs[j][i]) != firstStringRune {
				match = false
				break
			}

		}

		if match {
			commonPrefix += firstStringRune
		} else {
			break
		}
	}

	return commonPrefix
}
