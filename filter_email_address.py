import re

n = int(raw_input())
mails = list(raw_input() for _ in range(n))

def valid_email(mail):
    return re.match("^[a-z0-9_-]+@[a-z0-9]+\.[a-z0-9]{1,3}$", mail)

print sorted(filter(valid_email, mails))
