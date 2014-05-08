n = int(raw_input())
students = dict()

for i in range(n):
    marks = raw_input().split()
    name = marks[0]
    students[name] = [float(mark) for mark in marks[1:]]

selected = raw_input()
print "%.2f" % (sum(students[selected]) / len(students[selected]))
