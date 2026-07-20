/* 
 Problem Link: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
*/

package main

func searchRange(nums []int, target int) []int {
    lower := lowerBound(nums, target)

    if lower == len(nums) || nums[lower] != target {
        return []int{-1, -1}
    }

    upper := upperBound(nums, target)

    return []int{lower, upper - 1}
}

func lowerBound(nums []int, target int) int {
    left, right := 0, len(nums)

    for left < right {
        mid := left + (right-left)/2

        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return left
}

func upperBound(nums []int, target int) int {
    left, right := 0, len(nums)

    for left < right {
        mid := left + (right-left)/2

        if nums[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }

    return left
}

/*
    Time complexity: O(log n)
    Space complexity: O(1)
*/
