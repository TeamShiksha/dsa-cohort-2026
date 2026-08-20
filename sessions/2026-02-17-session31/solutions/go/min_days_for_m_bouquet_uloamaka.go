package main
/* 
	Problem link: https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/
*/

func minDays(bloomDay []int, m int, k int) int {
    if m*k > len(bloomDay) {
        return -1
    }

	max := 0
    for _, val := range bloomDay {
        if val > max {
            max = val
        }
    }

    left, right := 0, max
    minDays := -1

    for left <= right {
        mid := left + (right-left) / 2

        if getNumOfBouquets(bloomDay, mid, k) >= m {
            minDays = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return minDays
}

func getNumOfBouquets(bloomDay []int, mid int, k int) int {
    numOfBouquets := 0
    cnt := 0
    for i := 0; i < len(bloomDay); i++ {
        if bloomDay[i] <= mid {
            cnt++
        } else {
            cnt = 0
        }

        if cnt == k {
            numOfBouquets++
            cnt = 0
        }
    } 
    return numOfBouquets
}

/* 
	Time Complexity: O(n log m), where n is the length of bloomDay and m is the maximum value in bloomDay
	Space Complexity: O(1)
*/
