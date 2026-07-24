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
List<Integer> list = new ArrayList<>();
int index=0;
    public BSTIterator(TreeNode root) {
        
        
        inorder(root,list);

    }
    
    public int next() {
        
        return list.get(index++);

    }
    
    public boolean hasNext() {
        
        if(index+1 <= list.size())
        return true;
        else 
        return false;
    }

    public TreeNode inorder(TreeNode node,List<Integer> list)
    {
        if(node == null)
        return null;

        inorder(node.left,list);
        list.add(node.val);
        inorder(node.right,list);

        return node;
    }
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * BSTIterator obj = new BSTIterator(root);
 * int param_1 = obj.next();
 * boolean param_2 = obj.hasNext();
 */