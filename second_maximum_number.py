m = int(raw_input())
s = set(int(x) for x in raw_input().split())
l = sorted(s, reverse=True)
print l[1]
