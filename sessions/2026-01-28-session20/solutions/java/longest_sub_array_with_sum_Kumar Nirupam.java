import java.util.HashMap;

class Solution {
    public int longestSubarray(int[] arr, int k) {
        // code here
        int prefixSum =0;
        int res =0;
        
        HashMap<Integer , Integer > map = new HashMap<>();
        
        for(int i=0;i<= arr.length-1;i++){
            
            // continuous sum calculate
            prefixSum = prefixSum + arr[i];
            
            
            
            // store it in map
            if(!map.containsKey(prefixSum)){
                map.put(prefixSum , i);
            }
            
            // check if prefixSum == k
            
            if(prefixSum == k){
                res = i+1;
            }
            
            
            // check if complement (prefix-k) exist in map
            
            if(map.containsKey((prefixSum - k))){
                
                res = Math.max(res,( i - map.get(prefixSum - k)));
            }
            
        }
        return res;
    }
}
