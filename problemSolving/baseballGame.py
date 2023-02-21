class Solution:
    def calPoints(self, operations: List[str]) -> int:
        record = []

        for op in operations:
            if op == "C":
                record.pop()
            elif op == "D":
                s = record[-1]
                record.append(s * 2)
            elif op == "+":
                if len(record) >= 2:
                    s1, s2 = record[-1], record[-2]
                    record.append(s1 + s2)
            else:
                record.append(int(op))
        
        return sum(record)

# runtime - O(n) where n represents the number of operations 