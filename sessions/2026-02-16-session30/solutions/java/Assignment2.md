class Solution {
    public int findMin(int[] nums) {
        int low = 0;
        int high = nums.length-1;
        int minimum_Elememt= Integer.MAX_VALUE;

        while(low <= high){
            int mid = (low+high)/2;

            if( nums[low] <= nums[high]){
               minimum_Elememt =Math.min(minimum_Elememt ,nums[low]);
                break;
            }

           

         if( nums[mid] >= nums[low]){
            minimum_Elememt =Math.min(minimum_Elememt ,nums[low]);
            low = mid +1;

         }else{
           high = mid -1 ;
           minimum_Elememt = Math.min(minimum_Elememt ,nums[mid]);
         }


        }
         return minimum_Elememt ;
        
    }

   
}