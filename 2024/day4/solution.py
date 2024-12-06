from typing import Dict, List, Tuple


def p1(input: str) -> int:
    with open(input) as content:
        parsed = parse_grid(content.read().strip())
        result = find_xmas(parsed)
        return result


def p2(input: str) -> int:
    with open(input) as content:
        parsed = parse_grid(content.read().strip())
        result = find_cross_mas(parsed)
        return result


def parse_grid(content: str) -> List[List[str]]:
    lines = [l for l in content.split("\n")]
    grid = []
    for l in lines:
        grid.append([char for char in l])
    return grid


def _build_coords(grid: List[List[str]]) -> Dict[Tuple[int], str]:
    coords = {}
    for x in range(len(grid)):
        for y in range(len(grid[0])):
            coords[(x, y)] = grid[y][x]
    return coords


def find_cross_mas(grid: List[List[str]]) -> int:
    coords = _build_coords(grid)
    found = 0
    for coord in coords:
        found += check_cross_mas(coords, coord)
    return found


def find_xmas(grid: List[List[str]]) -> int:
    coords = _build_coords(grid)

    found = 0
    for coord in coords:
        found += check_xmas(coords, coord)

    return found


def check_xmas(coords: Dict[Tuple[int], str], start: Tuple[int]) -> bool:
    if coords.get(start, None) != "X":
        return 0

    found = 0

    if (
        # xmas going ->
        coords[start] == "X"
        and coords.get((start[0] + 1, start[1]), None) == "M"
        and coords.get((start[0] + 2, start[1]), None) == "A"
        and coords.get((start[0] + 3, start[1]), None) == "S"
    ):
        found += 1
    if (
        # xmas going <- (i.e. samx)
        coords[start] == "X"
        and coords.get((start[0] - 1, start[1]), None) == "M"
        and coords.get((start[0] - 2, start[1]), None) == "A"
        and coords.get((start[0] - 3, start[1]), None) == "S"
    ):
        found += 1
    if (
        # xmas going down
        coords[start] == "X"
        and coords.get((start[0], start[1] + 1), None) == "M"
        and coords.get((start[0], start[1] + 2), None) == "A"
        and coords.get((start[0], start[1] + 3), None) == "S"
    ):
        found += 1
    if (
        # xmas going up
        coords[start] == "X"
        and coords.get((start[0], start[1] - 1), None) == "M"
        and coords.get((start[0], start[1] - 2), None) == "A"
        and coords.get((start[0], start[1] - 3), None) == "S"
    ):
        found += 1
    if (
        # xmas going right down
        coords[start] == "X"
        and coords.get((start[0] + 1, start[1] + 1), None) == "M"
        and coords.get((start[0] + 2, start[1] + 2), None) == "A"
        and coords.get((start[0] + 3, start[1] + 3), None) == "S"
    ):
        found += 1
    if (
        # xmas going left down
        coords[start] == "X"
        and coords.get((start[0] - 1, start[1] + 1), None) == "M"
        and coords.get((start[0] - 2, start[1] + 2), None) == "A"
        and coords.get((start[0] - 3, start[1] + 3), None) == "S"
    ):
        found += 1
    if (
        # xmas going right up
        coords[start] == "X"
        and coords.get((start[0] + 1, start[1] - 1), None) == "M"
        and coords.get((start[0] + 2, start[1] - 2), None) == "A"
        and coords.get((start[0] + 3, start[1] - 3), None) == "S"
    ):
        found += 1
    if (
        # xmas going left up
        coords[start] == "X"
        and coords.get((start[0] - 1, start[1] - 1), None) == "M"
        and coords.get((start[0] - 2, start[1] - 2), None) == "A"
        and coords.get((start[0] - 3, start[1] - 3), None) == "S"
    ):
        found += 1
    return found


def check_cross_mas(coords: Dict[Tuple[int], str], start: Tuple[int]) -> bool:
    if coords.get(start, None) != "A":
        return 0

    # rule out cases where corners aren't S or M
    tl = (start[0] - 1, start[1] - 1)
    tr = (start[0] + 1, start[1] - 1)
    bl = (start[0] - 1, start[1] + 1)
    br = (start[0] + 1, start[1] + 1)

    tl_corner = coords.get(tl, None)
    tr_corner = coords.get(tr, None)
    bl_corner = coords.get(bl, None)
    br_corner = coords.get(br, None)

    for corner in (tl_corner, tr_corner, bl_corner, br_corner):
        if corner not in ("S", "M"):
            return 0

    found = 0
    if tl_corner == "S" and br_corner == "M":
        if (tr_corner == "S" and bl_corner == "M") or (
            tr_corner == "M" and bl_corner == "S"
        ):
            found += 1
    if tl_corner == "M" and br_corner == "S":
        if (tr_corner == "S" and bl_corner == "M") or (
            tr_corner == "M" and bl_corner == "S"
        ):
            found += 1

    return found


if __name__ == "__main__":
    part1 = p1("input.txt")
    print("Part1: {}".format(part1))
    part2 = p2("input.txt")
    print("Part2: {}".format(part2))
