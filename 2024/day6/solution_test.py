from textwrap import dedent
from .solution import *
import pytest


FIXME = None


def test_grid_from_str():
    # fmt: off
    input = dedent(
    """
    ....#.....
    .........#
    ..........
    ..#.......
    .......#..
    ..........
    .#..^.....
    ........#.
    #.........
    ......#...
    """
    )
    # fmt: on
    got = Grid.from_str(input)
    assert len(got.visited) == 0
    assert got.coords[(0, 0)] == "."
    assert got.coords[(4, 0)] == "#"
    assert got.start_pos == (4, 6)
    assert got.dir == "u"
    assert got.current_pos == (4, 6)


def test_e2e_p1():
    got = p1("example.txt")
    want = 41
    assert got == want


def test_e2e_p2():
    got = p2("example.txt")
    want = FIXME
    assert want == got
