// package main

// import "fmt"

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

// func main() {
// 	nums1 := []int{1, 3, 5, 6}
// 	target1 := 5
// 	fmt.Println(searchInsert(nums1, target1))

// 	nums2 := []int{1, 3, 5, 6}
// 	target2 := 2
// 	fmt.Println(searchInsert(nums2, target2))

// 	nums3 := []int{1, 3, 5, 6}
// 	target3 := 7
// 	fmt.Println(searchInsert(nums3, target3))
// }
