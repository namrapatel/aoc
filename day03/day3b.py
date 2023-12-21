grid = open("input.txt").read().splitlines()
rows, cols = len(grid), len(grid[0])
dirs = [[0,1],[1,0],[-1,0],[0,-1],[1,-1],[-1,1],[1,1],[-1,-1]]
total = 0

for r, row in enumerate(grid):
    for c, char in enumerate(row):
        if char != "*":
            continue

        starting_coords = set()

        # found an asterisk, move all in dirs to find a number
        for dr, dc in dirs:
            new_row, new_col = r+dr, c+dc
            if new_row < 0 or new_col < 0 or new_row >= rows or new_col >= cols or not grid[new_row][new_col].isdigit():
                continue
            # found a number, move to it's first index
            while new_col > 0 and grid[new_row][new_col-1].isdigit():
                new_col -= 1
            # found starting index, add to set
            starting_coords.add((new_row, new_col))

       # ignore if 2 numbers were not found 
        if len(starting_coords) != 2:
            continue

        # get the full number, multiply and add to total
        numbers = []
        for cr, cc in starting_coords:
            number = ""
            while cc < cols and grid[cr][cc].isdigit():
                number += grid[cr][cc]
                cc += 1
            numbers.append(int(number))

        total += numbers[0] * numbers[1]
print(total)