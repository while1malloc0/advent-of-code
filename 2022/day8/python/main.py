import argparse
from os import path
from typing import List


def parse_input(input: str) -> List[List[int]]:
    out = []
    for line in input.strip().split("\n"):
        row = []
        for item in line:
            row.append(int(item))
        out.append(row)
    return out


def is_visible(grid: List[List[int]], x: int, y: int) -> bool:
    if is_edge(grid, x, y):
        return True
    l = left(grid, x, y)
    r = right(grid, x, y)
    t = above(grid, x, y)
    b = below(grid, x, y)

    return (
        is_max(l, len(l) - 1) or is_max(r, 0) or is_max(t, len(t) - 1) or is_max(b, 0)
    )


def is_edge(grid: List[List[int]], x: int, y: int) -> bool:
    return x == 0 or x == len(grid[0]) - 1 or y == 0 or y == len(grid) - 1


def is_max(input: List[int], pos: int) -> bool:
    partitioned = []
    for (i, item) in enumerate(input):
        if i == pos:
            continue
        partitioned.append(item)
    return input[pos] == max(partitioned)


def below(grid: List[List[int]], x: int, y: int) -> List[int]:
    out = []
    for i in range(0, y + 1):
        out.append(grid[i][x])
    return out


def above(grid: List[List[int]], x: int, y: int) -> List[int]:
    out = []
    for i in range(y, len(grid)):
        out.append(grid[i][x])
    return out


def left(grid: List[List[int]], x: int, y: int) -> List[int]:
    out = []
    for i in range(0, x + 1):
        out.append(grid[y][i])
    return out


def right(grid: List[List[int]], x: int, y: int) -> List[int]:
    out = []
    for i in range(x, len(grid[0])):
        out.append(grid[y][i])
    return out


def part_one(input):
    # parsing txt -> matrix
    matrix = parse_input(input)
    num_visible = 0
    for i in range(0, len(matrix)):
        for j in range(0, len(matrix[0])):
            if is_visible(matrix, j, i):
                num_visible += 1
    return num_visible


def part_two(input):
    return input


parser = argparse.ArgumentParser(prog="aoc")
parser.add_argument("part", type=int)
parser.add_argument("--example", type=bool, default=False)


def main():
    args = parser.parse_args()
    if args.part not in (1, 2):
        raise TypeError("part argument must be one of: 1 or 2")

    in_file_name = "example.txt" if args.example else "input.txt"
    in_file_path = path.abspath(__file__ + f"/../../data/{in_file_name}")
    input = ""
    with open(in_file_path) as f:
        input = f.read()

    result = ""
    if args.part == 1:
        result = part_one(input)
    elif args.part == 2:
        result = part_two(input)
    else:
        raise Exception("unreachable")

    print(result)


if __name__ == "__main__":
    main()
