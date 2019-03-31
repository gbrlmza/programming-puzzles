#!/usr/bin/python3
'''
# Problem: https://code.google.com/codejam/contest/dashboard?c=32003#s=p1
'''

def move(x, y, direction):
    movments = {0:[1, 0], 90:[0, 1], 180:[-1, 0], 270:[0, -1]} # x,y movement per direction
    direction = (direction % 360)
    x += movments[direction][0]
    y += movments[direction][1]
    return [x, y]


def set_maze_coordinate(x, y, direction, room_steps, maze):
    if x not in maze:
        maze[x] = {}
    if y not in maze[x]:
        maze[x][y] = 0 # unvisited room
    # values for right, front, left and back, for direction 0(0º, facing east). 1=open, 0=close
    steps = {'W':'1010', 'WL':'0110', 'WR':'0011', 'WRR':'0010'}
    direction = round((direction % 360) / 90) # Simplify direction to 0-3(0º,90º,180º,270º)
    room_config = steps[room_steps]
    # rotate values to match current direction
    room_config = room_config[-direction:] + room_config[:len(room_config)-direction]
    # convert ENWS to EWSN for hex mazeping
    room_config = room_config[0:1] + room_config[2:3] + room_config[-1] + room_config[1:2]
    # OR bitwise to mix room configs from different path
    maze[x][y] = maze[x][y] | int(room_config, 2)


def navigate(x, y, direction, path, maze={}):
    room_steps = ''
    entry_direction = direction

    for step in path:
        if step == 'W':
            if room_steps:
                set_maze_coordinate(x, y, entry_direction, room_steps, maze)
                room_steps = ''
            entry_direction = direction
            x, y = move(x, y, direction)
        elif step == 'L':
            direction += 90
        elif step == 'R':
            direction -= 90

        room_steps += step

    return [x, y, direction]


def print_maze(maze):
    left = min(maze.keys())
    right = max(maze.keys())
    top = max(maze[left].keys())
    bottom = min(maze[left].keys())

    y = top
    while y >= bottom:
        row = ''
        x = left
        while x <= right:
            row += hex(maze[x][y])[2:]
            x += 1
        print(row)
        y -= 1


def main():
    cases = int(input())  # read a line with a single integer
    for i in range(1, cases + 1): # for each data line
        entry_to_exit, exit_to_entry = [s for s in input().split(" ")]
        maze = {}
        direction = 270
        x, y, direction = navigate(0, 1, 270, entry_to_exit, maze)
        x, y, direction = navigate(x, y, direction + 180, exit_to_entry, maze)
        print("Case #{}:".format(i))
        print_maze(maze)

main()