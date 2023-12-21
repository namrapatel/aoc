with open("input.txt", "r") as file:
    input = [line.strip() for line in file]

dirs = [[0,1],[1,0],[-1,0],[0,-1],[1,-1],[-1,1],[1,1],[-1,-1]]
rows, cols = len(input), len(input[0])
visited = [[False for _ in row] for row in input]


def search(input, row: int, col: int):
    num_found = 0
    nums = []
    # Move all all directions
    for dr, dc in dirs:
        new_row, new_col = row + dr, col + dc
        
        # Check to be within boundaries
        if new_row < 0 or new_col < 0 or new_row >= rows or new_col >= cols:
            continue

        # If char at next position is a number, grab the full number and append
        next = input[new_row][new_col]
        if next.isdigit() and (visited[new_row][new_col] == False): 
            full_num = grab_full_num(input, new_row, new_col)
            nums.append(full_num)
            num_found += 1
            if num_found == 2:
                return True, nums

    return False, nums
            
            
def grab_full_num(input, row: int, col: int):
    # Move to the front of the number
    while col > 0 and input[row][col-1].isdigit():
        col -= 1

    # Build the full number from the front and return
    full_num = ""
    while col < cols and (input[row][col].isdigit()):
        full_num = full_num + input[row][col]
        visited[row][col] = True
        col += 1
    return full_num


sum = 0
currNum = ""
for row in range(rows):
    for col in range(cols):
        curr = input[row][col]
        if curr == '*':
            valid, nums = search(input, row, col)
            if valid:
                new_sum = int(nums[0]) * int(nums[1])
                sum += new_sum

print(sum)