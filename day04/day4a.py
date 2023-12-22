total = 0
for x in open("test_input.txt"):
    # Remove "Card X:" prefix and strip whitespace
    x = x.split(":")[1].strip()

    # Get winning nums (a) and our nums (b) in lists of int
    # Note: k.split() turns str into list
    a,b = [list(map(int, k.split())) for k in x.split(" | ")]

    # For items in b, if item in a, then True (same as 1 in Python)
    intersection = sum(q in a for q in b)
    if intersection > 0:
        total += 2 ** (intersection-1)

print(total)