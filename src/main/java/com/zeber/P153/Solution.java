package com.zeber.P153;

class Solution {
    public int findMin(int[] nums) {
        int left = 0;
        int right = nums.length - 1;
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] > nums[right]) {
                left = mid + 1;
            } else if (nums[mid] < nums[right]) {
                right = mid;
            }
        }
        return nums[left];
    }

    //  1 2 3 4 5 6 7 8 9
    //   7 8 9 1 2 3 4 5 6
    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.findMin(new int[]{3,4,5,1,2}));
    }
}
