class Solution:
    def calPoints(self, operations: List[str]) -> int:
        record = []

        for op in operations:
            if op == "C":
                record.pop()
            elif op == "D":
                record.append(record[-1] * 2)
            elif op == "+":
                if len(record) >= 2:
                    a, b = record[-1], record[-2]
                    record.append(a + b)
            else:
                record.append(int(op))
        
        return sum(record)
        