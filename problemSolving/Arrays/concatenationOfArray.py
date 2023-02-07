class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        arr = [0 for i in range(2 * len(nums))]
        n = len(nums)

        for i in range(len(nums)):
            arr[i] = nums[i]
            arr[i + n] = nums[i]
        
        return arr
        