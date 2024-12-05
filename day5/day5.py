from collections import defaultdict
def checkValid(nums):
    cur = set()
    for n in nums:
        for c in cur:
            if c in validPath[n]:
                return False

        cur.add(n)
    return True
        
def customSort(nums):
    res = [n for n in nums]
    for i in range(len(res)):
        for j in range(len(res)):
            if res[i] in validPath[res[j]]:
                res[j], res[i] = res[i], res[j] 
    return res

with open("./day5/input.txt", "r") as file:
    solve = False
    validPath = defaultdict(list)
    good = 0
    bad = 0

    for line in file:
        if len(line) == 1:
            solve = True
            continue
        if solve:
            nums = line.strip("\n").split(",")
            if checkValid(nums):
                good += int(nums[len(nums) // 2])
            else:
                validOrder = customSort(nums)
                print(nums)
                print(validOrder)
                print()
                bad += int(validOrder[len(validOrder) // 2])
        else:
            a, b = line[:-1].split("|")
            validPath[a].append(b)
    print(bad)