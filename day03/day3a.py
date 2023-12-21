grid = open("input.txt").read().splitlines()
starting_coords = set()
rows, cols = len(grid), len(grid[0])
dirs = [[0,1],[1,0],[-1,0],[0,-1],[1,-1],[-1,1],[1,1],[-1,-1]]

for r, row in enumerate(grid):
    for c, char in enumerate(row):
        if char.isdigit() or char == ".":
            continue
        # found a symbol, move all in dirs to find a number
        for dr, dc in dirs:
            new_row, new_col = r+dr, c+dc
            if new_row < 0 or new_col < 0 or new_row >= rows or new_col >= cols or not grid[new_row][new_col].isdigit():
                continue
            # found a number, move to it's first index
            while new_col > 0 and grid[new_row][new_col-1].isdigit():
                new_col -= 1
            # found starting index, add to set
            starting_coords.add((new_row, new_col))

# get the full number, add to list of numbers and sum
numbers = []
for r, c in starting_coords:
    number = ""
    while c < cols and grid[r][c].isdigit():
        number += grid[r][c]
        c += 1
    numbers.append(int(number))

print(sum(numbers))