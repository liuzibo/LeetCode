package com.zeber.P31;

import java.util.Arrays;

public class Solution {
    public void nextPermutation(int[] nums) {
        int len = nums.length;
        int i = len - 2;
        while (i >= 0 && nums[i] >= nums[i + 1]) {
            i--;
        }
        int j = len - 1;
        if (i >= 0) {
            while (j >= 0 && nums[j] <= nums[i]) {
                j--;
            }
            swap(nums, i, j);
        }
        int left = i + 1;
        int right = len - 1;
        while (left < right) {
            swap(nums, left, right);
            left++;
            right--;
        }
// 1 5 8 4 7 6 5 3 1
// 1 5 8 5 1 3 4 6 7
    }

    public void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        int[] arr = new int[]{1, 1};
        solution.nextPermutation(arr);
        System.out.println(Arrays.toString(arr));
    }

}
