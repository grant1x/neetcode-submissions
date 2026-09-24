class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        maxLen = 0
        sett = set()

        leftIndex = 0
        for rightIndex in range(len(s)):

            while s[rightIndex] in sett:
                sett.remove(s[leftIndex])
                leftIndex = leftIndex + 1

            sett.add(s[rightIndex])

            maxLen = max(maxLen, rightIndex - leftIndex + 1)
          
        return maxLen
