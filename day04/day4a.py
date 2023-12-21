sum = 0

with open("input.txt") as file:
    for line in file:
        line = line.strip().split(" ")[2:]
        line = [item for item in line if item != '']

        winning = set()
        ours = set()

        past_line = False
        for i in range(len(line)):    
            if line[i] == '|':
                past_line = True
                continue
            if past_line:
                ours.add(line[i])
            else:
                winning.add(line[i])

        num_matches = len(winning.intersection(ours))
        if num_matches == 0:
            continue
        sum += 2**(num_matches-1)

print(sum)