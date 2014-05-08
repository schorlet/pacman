m = int(raw_input())
ms = set(int(x) for x in raw_input().split())
n = int(raw_input())
ns = set(int(x) for x in raw_input().split())
diff = [str(x) for x in sorted(ms ^ ns)]
print '\n'.join(diff)
