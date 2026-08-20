/* 
	Problem Link: https://leetcode.com/problems/search-in-rotated-sorted-array/
*/
package main

func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            return mid
        }

		// Check if the right half is sorted
        if nums[mid] <= nums[right] {
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        } else {
            // The left half is sorted
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }
    }

    return -1
}

/* 
	Time Complexity: O(log n)
	Space Complexity: O(1)
*/