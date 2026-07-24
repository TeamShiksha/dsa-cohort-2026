/* 
    Problem Link: https://www.geeksforgeeks.org/problems/ceil-in-a-sorted-array/1
*/
class Solution {
  public:
    int findCeil(vector<int>& arr, int x) {
        int left = 0;
        int right = arr.size();
        
        while (left < right) {
            int mid = left + (right - left) / 2;
            
            if (arr[mid] >= x)
            {
                right = mid;
            }
            else 
            {
                left = mid+1;
            }
        }

        if (left == arr.size()) {
            return -1;
        }
        
        return left;
    }
};

/*
    Time complexity: O(log n)
    Space complexity: O(1)
*/
