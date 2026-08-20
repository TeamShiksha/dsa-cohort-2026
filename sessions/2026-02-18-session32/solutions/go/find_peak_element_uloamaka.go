/* 
	Probllem link: https://leetcode.com/problems/find-peak-element/
*/

package main

func findPeakElement(nums []int) int {
    left, right := 0, len(nums)-1

    for left < right {
        mid := left + (right-left)/2

        if nums[mid] > nums[mid+1] {
            right = mid
        } else {
            left = mid + 1
        }
    }

    return left
}

/* 
	Time complexity: O(log n),
	Space complexity: O(1)
*/