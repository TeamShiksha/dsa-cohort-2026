#include <iostream>
#include <vector>
#include <algorithm>
#include <limits>
#include <climits>

using namespace std;

class Solution {
public:
    // helper function to find the min and max in an array
    vector<int> findMinMax(vector<int>& bloomDay) {
        int minEle = INT_MAX;
        int maxEle = INT_MIN;

        for(int i = 0; i < bloomDay.size(); i++) {
            if(bloomDay[i] > maxEle) {
                maxEle = bloomDay[i];
            }
            if(bloomDay[i] < minEle) {
                minEle = bloomDay[i];
            }
        }

        return {minEle, maxEle};
    }


    // Greedy helper function 
    bool canMakeBouquets(vector<int>& bloomDay, int m, int k, int currentDay) {
        int bouquets = 0;
        int consecutive_flowers = 0;

        for (int day : bloomDay) {
            if (day <= currentDay) {
                consecutive_flowers++;
                if (consecutive_flowers == k) {
                    bouquets++;
                    consecutive_flowers = 0; // Greedily reset
                }
            } else {
                consecutive_flowers = 0;
            }
        }
        return bouquets >= m;
    }
    
    // brute force: greedy + linear search
    int minDaysBrute(vector<int>& bloomDay, int m, int k) {
        long long flowers = (long long)m * k;
        if(flowers > bloomDay.size()) {
            return -1;
        }

        vector<int> minMax = findMinMax(bloomDay); // ans is in form of {min, max}
        int minEle = minMax[0];
        int maxEle = minMax[1];
        

        // brute force: Linear Search 
        // We check every single day from minEle to maxEle
        for(int day = minEle; day <= maxEle; day++) {
            
            // The first day that works is guaranteed to be the minimum day
            if(canMakeBouquets(bloomDay, m, k, day)) {
                return day; 
            }
        }
        
        return -1;

        /*
            TC: O(n*(max-min)): due to nested loops (outer loop from 0 to n and inner loop from min to max) 
            SC: O(1): no extra space used

            NOTE: Gives TLE for large input sizes
        */
    }

    // optimal solution: greedy + binary search
    int minDaysOpt(vector<int>& bloomDay, int m, int k) {
        // m = number of bouquets
        // k = no. of flowers in each bouquet
        // each bloomDay[i] corresponds to a flower and when it will bloom

        // bouquets can only be made from adjacent flowers
        // if the number of flowers required: m*k > total flowers, we return -1
        long long flowers = (long long)m * k;
        if(flowers > bloomDay.size()) {
            return -1;
        }

        // binary search on answers: the search space is min day and max day
        // first we find the minimum and maximum blooming time
        vector<int> minMax = findMinMax(bloomDay); // ans is in form of {min, max}

        // since the search space is the range of blooming days
        // we initialize our binary search as such
        int low = minMax[0];
        int high = minMax[1];

        while(low < high) {
            int mid = low + (high - low)/2;

            if(canMakeBouquets(bloomDay, m, k, mid)) {
                high = mid;
            }
            else {
                low = mid + 1;
            }
        }
        return low;
    }

    /*
        TC: O(n* log (max-min) + O(n)) : where n = size of bloomDay array, max = maxEle in bloomDay, min = minEle in bloomDay
        n because of outer loop form 0 to n-1, log(max-min) because of binary search in the range [min, max] and O(n) to find minEle and maxEle
        SC: O(1): no extra space used
    */
};

int main() {
    vector<int> bloomDay = {7,7,7,7,12,7,7};
    int m = 2, k = 3;

    Solution s;

    int ans = s.minDaysOpt(bloomDay, m, k); // expected: 12
    cout<<ans<<endl;
}