import re
n = int(raw_input())
for _ in range(n):
    b = re.match("^[7-9][0-9]{9}$", raw_input())
    print "YES" if b else "NO"

