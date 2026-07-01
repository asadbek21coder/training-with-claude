package ds

import "unicode"

// TwoSumSorted finds two numbers in a sorted array that add up to target.
// Returns their indices (0-based). Array is guaranteed to have exactly one solution.
// Example: [1,3,5,7,9,11], target=14 → [1, 5] (3 + 11 = 14)
func TwoSumSorted(arr []int, target int) [2]int {
	left, right := 0, len(arr)-1

	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return [2]int{left, right}
		}
		if sum > target {
			right -= 1
		} else {
			left += 1
		}
	}

	return [2]int{}
}

// IsPalindrome checks if a string reads the same forwards and backwards.
// Ignore case, only consider letters and digits.
// Example: "A man a plan a canal Panama" → true
func IsPalindrome(s string) bool {
	clearedStr := []rune{}
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			clearedStr = append(clearedStr, v)
		}
	}
	left, right := 0, len(clearedStr)-1
	for left < right {
		if unicode.ToLower(clearedStr[left]) != unicode.ToLower(clearedStr[right]) {
			return false
		} 
		left += 1
		right -= 1
	}

	return true
}
