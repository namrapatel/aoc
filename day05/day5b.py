inputs, *blocks = open("input.txt").read().split("\n\n")
inputs = list(map(int, inputs.split(": ")[1].split()))
seeds = []

# populate seeds with (start, end) tuples from inputs
for i in range(0, len(inputs), 2):
    seeds.append((inputs[i], inputs[i]+inputs[i+1]))

# for each map
for block in blocks:
    ranges = []
    
    # get the numbers from the map as a list of ints
    for line in block.splitlines()[1:]:
        ranges.append(list(map(int, line.split())))
    
    new = []
    while len(seeds) > 0:
        s, e = seeds.pop()
        for a, b, c in ranges:
            overlap_start = max(s,b)
            overlap_end = min(e, b + c)
            if overlap_start < overlap_end:
                # append tuple representing the core overlap after performing shift defined in the map
                new.append((overlap_start - b + a, overlap_end - b + a))
                # check for leftover on left and right sides of overlap_start and overlap_end
                if overlap_start > s:
                    seeds.append((s, overlap_start))
                if e > overlap_end:
                    seeds.append((overlap_end, e))
                break
        else:
            new.append((s, e))

    seeds = new

print(min(seeds))