/*
Approach Used: Recursive Backtracking with Extra Space

1. The reverseArray() function checks if the array size is less than 2.
   - If yes, no reversal is needed and we return immediately.

2. We create an ArrayList 'res' to temporarily store reversed elements.

3. The reverse() function is a recursive helper method.
   - It traverses the array from index 0 to n using recursion.
   - Base Condition: When index == arr.length, recursion stops.
   - This ensures we do not access out-of-bounds index.

4. Recursive Flow:
   - First, recursion goes till the end of the array.
   - While returning from recursive calls (backtracking phase),
     elements are added to 'res'.
   - Since elements are added during the return phase,
     they are stored in reverse order.

5. Finally, we copy elements from 'res' back into the original array.

Time Complexity: O(n)
- Each element is visited once.

Space Complexity: O(n)
- O(n) for recursion call stack.
- O(n) for the extra ArrayList used to store reversed elements.
*/



import java.util.ArrayList;

class Solution {
    public void reverseArray(int arr[]) {
        // code here
        
        
        int n = arr.length;
        
        if(n<2){
            return ;
        }
        
        ArrayList<Integer> res = new ArrayList<>();
        reverse(arr , res , 0);
        
        int j =0;
        for(int i =0 ; i<res.size() ; i++){
            arr[j] = res.get(i);
            j++;
        }
        
        return  ;
    }
    
    private void reverse(int[] arr ,ArrayList<Integer> res ,int index){
        
        if(index == arr.length ){
           return;
        }
        
        reverse(arr , res , index+1);
        // Last Recursion 
         res.add(arr[index]);
    }
}