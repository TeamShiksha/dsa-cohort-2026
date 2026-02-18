#include <iostream>
#include <vector>
#include <algorithm>
#include <limits>

using namespace std;


class Solution {
public:
    // brute force solution using linear search
    bool searchBrute(vector<int>& nums, int target) {
        // brute force solution: binary search

        // edge cases
        if(nums[0] == target || nums[nums.size()-1] == target) {
            return true;
        }

        for(int i = 1; i < nums.size()-1; i++) {
            if(nums[i] == target) {
                return true;
            }
        }

        return false;

        // TC: O(n) : linear search checks every element => O(n) for array containing n elements
        // SC: O(1) : No extra space used
    }


    // optimal solution using binary search
    bool searchOpt(vector<int>& nums, int target) {
        // optimal solution: binary search because at either side of pivot element, the search space is sorted

        // if first or the last element is the target
        if(nums[0] == target || nums[nums.size()-1] == target) {
            return true;
        }

        int low = 0;
        int high = nums.size()-1;

        while(low <= high) {
            int mid = low + (high - low)/2;

            // checking the conditions
            // case 1: nums[mid] is the target
            if(nums[mid] == target) {
                return true;
            }

            // due to rotation and repetition, there is a 
            // chance that elements at both ends are equal to mid
            if((nums[mid] == nums[low]) && (nums[high] == nums[mid])) {
                low++;
                high--;
            }
            else if(nums[low] <= nums[mid]) {
                // first half is sorted => implement binary search
                if((nums[low] <= target) && (target < nums[mid])) {
                    // target is present in first half
                    high = mid - 1;
                }
                else {
                    // target is not present in first half
                    low = mid + 1;
                }
            }
            else {
                // second half is sorted => apply binary search
                if((nums[mid] < target) && (target <= nums[high])) {
                    // target is in 2nd half
                    low = mid + 1;
                }
                else {
                    // target is not in 2nd half
                    high = mid - 1;
                }
            }
        } 

        return false;

        /*
            TC: O(log n): binary search halves the search space at each iteration hence log2(n) number of operations
            SC: O(1): no extra space has been used
        */
    }
};


int main() {
        vector<int> nums = {2,5,6,0,0,1,2};
        int target1 = 3;
        int target2 = 0;

        Solution s;
        bool ans = s.searchOpt(nums, target1);
        cout<<ans<<endl; // expected 0 => false


        bool ans2 = s.searchOpt(nums, target2);
        cout<<ans2<<endl; // expected 1 => True
    return 0;
}