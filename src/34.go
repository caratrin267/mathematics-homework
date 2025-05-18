class Solution {
    public int findMin(int[] nums) {
        if (nums == null || nums.length <= 1) return 0;
        int left = 0;
        int right = nums.length - 1;
        
        while (left < right) {
            int mid = left + (right - left) / 2;
            
            // If the middle element is greater than its successor,
            if (nums[mid] > nums[mid + 1]) return mid;
            else if (nums[mid] < nums[left]) {
                right = mid;
            } else {
                left = mid + 1;
            }
        }

        return -1; // If the loop completes without returning, there are no pairs to find.
    }
}
