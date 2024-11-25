package main

import (
	"fmt"
)

func CollapseStrings(input []string) map[string]string {
	prefixMap := make(map[string][]string)
	// Build map from prefix of length 6 to strings
	for _, s := range input {
		if len(s) < 6 {
			// Map to itself
			prefixMap[s] = append(prefixMap[s], s)
			continue
		}
		prefix6 := s[:6]
		prefixMap[prefix6] = append(prefixMap[prefix6], s)
	}

	result := make(map[string]string)

	for _, group := range prefixMap {
		if len(group) >= 3 {
			// Find the longest common prefix
			lcp := longestCommonPrefix(group)
			if len(lcp) < 6 {
				lcp = group[0][:6]
			}
			groupName := fmt.Sprintf("\"%s(%d)\"", lcp, len(group))
			for _, s := range group {
				result[s] = groupName
			}
		} else {
			// Map to itself
			for _, s := range group {
				result[s] = s
			}
		}
	}

	return result
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = commonPrefix(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func commonPrefix(s1, s2 string) string {
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}
	i := 0
	for i < minLen && s1[i] == s2[i] {
		i++
	}
	return s1[:i]
}
