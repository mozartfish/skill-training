class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        parens = {")": "(", "}": "{", "]":"["}

        for ch in s:
            if ch not in parens:
                stack.append(ch)
            elif not stack or stack[-1] != parens[ch]:
                return False 
            else:
                stack.pop()
            
        return not stack 
# runtime - O(n) where n represents the size of the string