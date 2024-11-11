package main

func GroupStrings(strings []string, commonCount int) map[string][]string {
	result := make(map[string][]string)
	longStrings := []string{}
	for _, s := range strings {
		if len(s) <= commonCount {
			result[s] = []string{s}
		} else {
			longStrings = append(longStrings, s)
		}
	}
	return result
}
