first = raw_input().strip()
second = raw_input().strip()
print sum(1 for i in range(len(first) - len(second) + 1)
        if first[i:i+len(second)] == second)
