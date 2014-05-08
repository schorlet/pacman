def formats(phone):
    if len(phone) == 10: # no prefix
        return "+91 %s %s" % (phone[:5], phone[5:])
    elif len(phone) == 11: # '0' prefix
        return "+91 %s %s" % (phone[1:6], phone[6:])
    elif len(phone) == 12: # '91' prefix
        return "+91 %s %s" % (phone[2:7], phone[7:])
    elif len(phone) == 13: # '+91' prefix
        return "+91 %s %s" % (phone[3:8], phone[8:])


def standardizes(func):
    def inner(phones):
        func([formats(phone) for phone in phones])
    return inner


@standardizes
def print_sorted(phones):
    print "\n".join(sorted(phones))


n = int(raw_input())
nums = [raw_input().strip() for _ in range(n)]
print_sorted(nums)
