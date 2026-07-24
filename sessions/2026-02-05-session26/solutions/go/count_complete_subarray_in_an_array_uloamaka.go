/* 
	Problem link: https://leetcode.com/problems/count-complete-subarrays-in-an-array/
*/
package main

func countCompleteSubarrays(nums []int) int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }

    seen := make(map[int]int)

    left:= 0
    cnt := 0
    for right := 0; right < len(nums); right++ {
        seen[nums[right]]++

        for len(seen) == len(freq) {
            seen[nums[left]]--

            if seen[nums[left]] == 0 {
                delete(seen, nums[left])
            }

            left++
        }

        cnt += left
    }
    return cnt
}

/* 
	Time complexity: O(n)
    Space complexity: O(k)
*/