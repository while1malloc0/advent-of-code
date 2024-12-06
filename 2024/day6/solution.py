from typing import Dict, Set, Tuple
from copy import deepcopy


class Grid:
    visited: Set[Tuple[int]]

    @staticmethod
    def from_str(s: str):
        coords = dict()
        lines = [l.strip() for l in s.split("\n")]
        twod = []
        for line in lines:
            if line == "":
                continue
            chars = [c for c in line]
            twod.append(chars)

        g = Grid()

        for x in range(len(twod)):
            for y in range(len(twod[0])):
                c = twod[y][x]
                if c == "^":
                    g.start_pos = (x, y)
                    g.current_pos = g.start_pos
                coords[(x, y)] = c

        g.coords = coords
        g.dir = "u"
        g.current_loop = set()
        g.last_loop = set()
        g.size_x = len(twod)
        g.size_y = len(twod[0])

        return g

    def __init__(
        self, coords: Dict[Tuple[int], str] = None, start_pos: Tuple[int] = (0, 0)
    ):
        self.visited = set()
        self.coords = coords
        self.start_pos = start_pos
        self.current_pos = start_pos
        self.dir = "u"
        self.last_loop = None
        self.current_loop = set()
        self.looping = False
        self.size_x = 0
        self.size_y = 0

    def play(self):
        while self.tick():
            pass
        return len(self.visited)

    def tick(self):
        # we're looping
        next_pos = self._next_pos()

        next_space = self.coords.get(next_pos, None)
        # we've run off the board
        if not next_space:
            return False

        while next_space == "#":
            self._rotate()
            next_pos = self._next_pos()
            next_space = self.coords.get(next_pos, None)
            if not next_space:
                return False

        print("c {}".format(self.current_loop))
        print("l {}".format(self.last_loop))
        if self.current_loop and self.last_loop and self.current_loop == self.last_loop:
            self.looping = True
            return False

        self.current_loop.add(self.current_pos)
        self.current_pos = next_pos
        self.visited.add(self.current_pos)

        return True

    def _next_pos(self):
        x, y = self.current_pos
        if self.dir == "u":
            return x, y - 1
        if self.dir == "d":
            return x, y + 1
        if self.dir == "l":
            return x - 1, y
        if self.dir == "r":
            return x + 1, y
        raise RuntimeError("_next_pos: we're traveling in an unknown direction")

    def _rotate(self):
        if self.dir == "u":
            self.dir = "r"
        elif self.dir == "d":
            self.dir = "l"
        elif self.dir == "l":
            self.dir = "u"
            self.last_loop = deepcopy(self.current_loop)
            self.current_loop.clear()
        elif self.dir == "r":
            self.dir = "d"
        else:
            raise RuntimeError(
                "_rotate: we're traveling in an unknown direction %s" % self.dir
            )


def p1(input: str) -> int:
    with open(input) as content:
        grid = Grid.from_str(content.read().strip())
        result = grid.play()
        return result


def p2(input: str) -> int:
    with open(input) as content:
        loops = 0
        grid = Grid.from_str(content.read().strip())
        for x in range(grid.size_x):
            for y in range(grid.size_y):
                cp = deepcopy(grid)
                if cp.coords.get((x, y)) != "#":
                    cp.coords[(x, y)] = "#"
                    cp.play()
                    if cp.looping:
                        loops += 1
        return loops


if __name__ == "__main__":
    part1 = p1("input.txt")
    print("Part1: {}".format(part1))
    part2 = p2("input.txt")
    print("Part2: {}".format(part2))
