/* 
	Problem link: https://leetcode.com/problems/sqrtx/
*/
package main

func mySqrt(x int) int {
    left, right := 0, x 
    
    if x < 2 {
        return x
    }

    for left < right {
        mid := left + (right-left) / 2

        if mid > x/mid { // same as mid*mid > x, but avoids overflow
            right = mid
        } else {
            left = mid+1
        }
    }
    return left-1
}

/* 
	Time Complexity: O(log n)
	Space Complexity: O(1)
*/
