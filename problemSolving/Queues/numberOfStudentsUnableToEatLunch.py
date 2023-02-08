class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        from collections import deque
        stuq = deque(students)
        sandq = deque(sandwiches)

        while stuq and sandq and sandq[0] in stuq:
            if stuq[0] == sandq[0]:
                stuq.popleft()
                sandq.popleft()
            else:
                student = stuq.popleft()
                stuq.append(student)
        
        return len(stuq)
        