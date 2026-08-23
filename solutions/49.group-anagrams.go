package solutions

import "slices"

// @leet start
func stringSort(s string) string {
	byteArray := []byte(s)
	slices.Sort(byteArray)
	return string(byteArray)
}

func groupAnagrams(strs []string) [][]string {
	aas := make(map[string][]string)

	for _, currString := range strs {
		byteSum := stringSort(currString)
		currGrp := aas[byteSum]
		aas[byteSum] = append(currGrp, currString)
	}

	var anagramGroup [][]string

	for _, group := range aas {
		anagramGroup = append(anagramGroup, group)
	}

	return anagramGroup
}

// @leet end
