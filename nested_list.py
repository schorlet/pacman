from operator import itemgetter

n = int(raw_input())
students = list()
for i in range(n):
    name = raw_input()
    marks = float(raw_input())
    students.append((name, marks))

# sort by name
students.sort(key=itemgetter(0))
# sort by marks
students.sort(key=itemgetter(1))

lowest = students[0][1]
lower = lowest

for i in range(1, n):
    mark = students[i][1]
    if mark == lowest:
        continue
    name = students[i][0]
    if lowest == lower:
        lower = mark
        print name
    elif mark == lower:
        print name
    else:
        break
