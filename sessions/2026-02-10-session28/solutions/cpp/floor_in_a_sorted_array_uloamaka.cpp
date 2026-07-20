/* 
    Problem Link: https://www.geeksforgeeks.org/problems/floor-in-a-sorted-array-1587115620/1
*/
class Solution {
  public:
    int findFloor(vector<int>& arr, int x) {
        int left = 0;
        int right = arr.size();
        
        while (left < right) {
            int mid = left + (right-left) / 2;
            
            if (arr[mid] > x) 
            {
                right = mid;
            } 
            else 
            {
                left = mid + 1;
            }
        }

        if (right == 0) {
            return -1;
        }

        return left-1;
    }
};

/* 
    Time complexity: O(log n)
    Space complexity: O(1)
*/
