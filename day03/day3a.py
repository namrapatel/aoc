with open("input.txt", "r") as file:
    input = [line.strip() for line in file]

dirs = [[0,1],[1,0],[-1,0],[0,-1],[1,-1],[-1,1],[1,1],[-1,-1]]
rows, cols = len(input), len(input[0])
visited = [[False for _ in row] for row in input]


def search(input, row: int, col: int):
    # Move all all directions
    for dr, dc in dirs:

        new_row, new_col = row + dr, col + dc
        # Check to be within boundaries
        if new_row < 0 or new_col < 0 or new_row >= rows or new_col >= cols:
            continue

        # If char at next dir is symbol mark origin as visited and return True
        next = input[new_row][new_col]
        if not (next.isdigit()) and not (next == '.'):
            visited[row][col] = True
            return True

    return False
            
            
def grabFullNum(input, row: int, col: int, currNum: str):
    while col + 1 < cols and (input[row][col+1].isdigit()):
        currNum = currNum + input[row][col+1]
        visited[row][col+1] = True
        col += 1
    return currNum


sum = 0
currNum = ""
for row in range(rows):
    for col in range(cols):
        curr = input[row][col]
        if curr.isdigit() and (visited[row][col] == False):
            currNum = currNum + curr
            if search(input, row, col):
                num = grabFullNum(input, row, col, currNum)
                sum = sum + int(num)
                currNum = ""
            elif col + 1 >= cols or not (input[row][col+1].isdigit()) or (visited[row][col+1] == True):
                currNum = ""


print(sum)