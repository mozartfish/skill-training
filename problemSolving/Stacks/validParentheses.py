class Solution:
    def isValid(self, s: str) -> bool:
        parens = {")" : "(", "}" : "{", "]" : "["}
        stack = []
        for p in s:
            if p not in parens:
                stack.append(p)
            elif not stack or stack[-1] != parens[p]:
                return False
            else:
                stack.pop() 
             
        return not stack

        