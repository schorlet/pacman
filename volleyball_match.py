import math

a = long(raw_input())
b = long(raw_input())

c = a + b - 1
d = min(a, b)
e = max(a, b)

# print a, b, c, d, e

if e < 25:
    print 0
elif e == 25 and d == 24:
    print 0
elif e > 25 and e - d != 2:
    print 0
else:
    f = 10**9 + 7
    g = long(math.factorial(c) / (math.factorial(d) * math.factorial(c-d)))
    if g > f:
        print g % f
    else:
        print g



print
# combinations
print math.factorial(e) / (math.factorial(d) * math.factorial(e-d))

# permutations
print math.factorial(e) / math.factorial(e-d)
