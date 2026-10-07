package ds

// BinarySearch finds target in a sorted array.
// Returns the index of target, or -1 if not found.
func BinarySearch(arr []int, target int) int {
	left, right, mid := 0, len(arr)-1,0

	for left <= right {
		mid = left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}

// SearchInsertPosition finds target in sorted array.
// If not found, returns the index where it should be inserted.
// Example: [1,3,5,7], target=4 → return 2 (insert between 3 and 5)
func SearchInsertPosition(arr []int, target int) int {
	left, right, mid := 0, len(arr)-1, 0
	for left <= right {
		mid = left + (right-left)/2
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
