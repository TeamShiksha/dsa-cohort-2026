/* 
	Problem link: https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
*/

package main

func search2(nums []int, target int) bool {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left) / 2

        if nums[mid] == target {
            return true
        }

        // Cannot determine which half is sorted, decrement both sides
        if nums[left] == nums[mid] && nums[mid] == nums[right] {
            left++
            right--
            continue
        }

        if nums[left] <= nums[mid] {
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }

    return false
}

/* 
	Time complexity: O(log n) in average case, O(n) in worst case
	Space complexity: O(1)

*/