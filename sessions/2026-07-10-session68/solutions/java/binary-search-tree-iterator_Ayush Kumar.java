import TreeNode;

package solutions.java;
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class BSTIterator {

    Stack<TreeNode> stack = new Stack<>();

    public BSTIterator(TreeNode root) {
        
        TreeNode node = root; 
        while(node != null)
        {
            stack.push(node);
            node = node.left;
        }
    }
    
    public int next() {
        

        TreeNode smaller = stack.pop();
        
        if(smaller.right != null)
        {
            //stack.push(smaller.right);
            TreeNode node = smaller.right;

            while(node != null)
            {
                stack.push(node);
                node = node.left;
            }
        }

        return smaller.val;
    }
    
    public boolean hasNext() {

        if(stack.isEmpty())
        return false;
        else
        return true;
        
    }




}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * BSTIterator obj = new BSTIterator(root);
 * int param_1 = obj.next();
 * boolean param_2 = obj.hasNext();
 */