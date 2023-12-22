m = {}

for i, x in enumerate(open("test_input.txt")):
    # Card 1 will never be in m, account for it
    if i not in m:
        m[i] = 1

    # Remove "Card X:" prefix and strip whitespace
    x = x.split(":")[1].strip()

    # Get winning nums (a) and our nums (b) in lists of int
    # Note: k.split() turns str into list
    a,b = [list(map(int, k.split())) for k in x.split(" | ")]

    # For items in b, if item in a, then True (same as 1 in Python)
    matches = sum(q in a for q in b)

    # For curr + 1 to num of matches + 1, add a copy
    for n in range(i + 1, i + matches + 1):
        m[n] = m.get(n, 1) + m[i]
    
# sum all the copies
print(sum(m.values())) 