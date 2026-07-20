/* 
	Problem Link: https://leetcode.com/problems/search-a-2d-matrix-ii/
*/
package main

func searchMatrix(matrix [][]int, target int) bool {
    for i := 0; i < len(matrix); i++ {
        left := 0
        right := len(matrix[i])

        for left < right {
            mid := left + (right - left ) / 2
            
            if matrix[i][mid] == target {
                return true
            }

            if matrix[i][mid] < target {
                left = mid+1
            } else {
                right = mid
            }
        }
    }
    return false
}

/* 
	Time complexity: O(m * log n),
	Space complexity: O(1)
*/
