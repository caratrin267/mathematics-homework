class Solution {
    public int[] nextGreaterElements(int[] nums) {
        int n = nums.length;
        
        // Initialize two arrays: one with results and another with indices of elements in nums.
        int[] result = new int[n];
        int[] index = new int[2 * n - 1];
        
        for (int i = 0; i < n; ++i) {
            if (nums[i] == 0) {
                nums[i] = -1;
            }
            else {
                result[index[2 * i + 1]] = nums[i];
                index[2 * i + 1]++;
            }
        }
        
        for (int j = n - 1; j >= 0; --j) {
            while (index[j] > 0 && nums[index[j]] < nums[j]) {
                result[index[index[j]]++] = nums[j];
                index[j]--;
            }
        }

        return Arrays.copyOf(result, n);
    }
}
