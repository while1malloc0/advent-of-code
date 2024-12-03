from .solution import *
from textwrap import dedent


def test_e2e_day1():
    got = p1("example.txt")
    want = 11
    assert got == want


def test_distances_from_cols():
    input = [[1, 2], [2, 3], [3, 4]]
    want = [1, 1, 1]
    got = distances_from_cols(input)
    assert want == got


def test_sort_pairs():
    input = [[3, 4], [4, 3], [2, 5], [1, 3], [3, 9], [3, 3]]
    want = [[1, 3], [2, 3], [3, 3], [3, 4], [3, 5], [4, 9]]
    got = sort_pairs(input)
    assert want == got


def test_parse_pairs():
    input = dedent(
        """
    3   4
    4   3
    2   5
    1   3
    3   9
    3   3
    """
    )
    got = parse_pairs(input)
    want = [[3, 4], [4, 3], [2, 5], [1, 3], [3, 9], [3, 3]]
    assert got == want


def test_calc_occurrances():
    input = [[3, 4], [4, 3], [2, 5], [1, 3], [3, 9], [3, 3]]
    got = calc_occurances(input)
    want = [9, 4, 0, 0, 9, 9]
    assert got == want


def test_e2e_day1_p2():
    input = "example.txt"
    want = 31
    got = p2(input)
    assert want == got
