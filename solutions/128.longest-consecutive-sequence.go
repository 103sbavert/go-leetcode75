package solutions

// @leet start
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	anchors := make(map[int]bool, len(nums))

	for _, curr := range nums {
		anchors[curr] = true
	}

	maxLen := 1
	for curr := range anchors {
		currLen := 1
		if !anchors[curr-1] {
			for anchors[curr+1] {
				currLen++
				curr++
			}
		}

		maxLen = max(currLen, maxLen)
	}

	return maxLen
}

// @leet end

