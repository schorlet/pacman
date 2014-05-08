n = int(raw_input())
a, b = 0, 1
nfib = [a, b]
for i in range(n-2):
    a, b = b, b + a
    nfib.append(b)

cube = lambda x: x**3
print list(map(cube, nfib[:n]))

