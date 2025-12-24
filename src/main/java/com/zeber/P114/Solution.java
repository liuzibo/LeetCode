package com.zeber.P114;

class Solution {
    public void flatten(TreeNode root) {
        TreeNode cur = root;
        while (cur != null) {
            if (cur.left != null){
                TreeNode next = cur.left;
                TreeNode pre = cur.left;
                while(pre.right!=null) {
                    pre = pre.right;
                }
                pre.right = cur.right;
                cur.right = next;
                cur.left = null;
            }
            cur = cur.right;
        }
    }


    public static void main(String[] args) {
        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(5);
        root.left.left = new TreeNode(3);
        root.left.right = new TreeNode(4);
        root.right.left = null;
        root.right.right = new TreeNode(6);


        Solution s = new Solution();
        s.flatten(root);
    }
}
