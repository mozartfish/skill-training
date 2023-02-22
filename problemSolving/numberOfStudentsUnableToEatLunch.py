class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        from collections import deque
        studQ = deque(students)
        sandQ = deque(sandwiches)
        
        while studQ and sandQ and sandQ[0] in studQ:
            if studQ[0] == sandQ[0]:
                studQ.popleft()
                sandQ.popleft()
            else:
                student = studQ.popleft()
                studQ.append(student)
        
        return len(studQ)


        # runtime - the number of sandwiches that match with sandwiches[0]