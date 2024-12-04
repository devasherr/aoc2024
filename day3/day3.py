import re

pattern = r"mul\(\d+,\d+\)|do\(\)|don't\(\)"
matches = re.findall(pattern, open("input.txt").read())

res = 0
allow = True

for match in matches:
    if match == "do()":
        allow = True
    elif match == "don't()":
        allow = False
    else:
        if allow:
            x, y = map(int, match[4:-1].split(","))
            res += (x * y)
print(res)