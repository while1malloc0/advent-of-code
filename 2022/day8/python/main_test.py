from .main import *


def test_parse_input():
    input = """
30373
25512
65332
33549
35390
"""
    want = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = parse_input(input)
    assert want == got


def test_is_visible():
    input = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = is_visible(input, 0, 0)
    want = True
    assert got == want


def test_is_visible_not_edge():
    input = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = is_visible(input, 1, 1)
    want = True
    assert got == want


def test_is_edge_bottom():
    input = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = is_edge(input, 3, 0)
    want = True
    assert got == want


def test_is_edge_not_edge():
    input = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = is_edge(input, 1, 1)
    want = False
    assert got == want


def test_is_edge_top():
    input = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = is_edge(input, 1, 4)
    want = True
    assert got == want


def test_is_edge_right():
    input = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = is_edge(input, 4, 3)
    want = True
    assert got == want


def test_is_edge_left():
    input = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = is_edge(input, 0, 3)
    want = True
    assert got == want


def test_below():
    input = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = below(input, 1, 1)
    want = [0, 5]
    assert got == want


def test_above():
    input = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = above(input, 1, 1)
    want = [5, 5, 3, 5]
    assert got == want


def test_left():
    input = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = left(input, 1, 1)
    want = [2, 5]
    assert got == want


def test_right():
    input = [
        [3, 0, 3, 7, 3],
        [2, 5, 5, 1, 2],
        [6, 5, 3, 3, 2],
        [3, 3, 5, 4, 9],
        [3, 5, 3, 9, 0],
    ]
    got = right(input, 1, 1)
    want = [5, 5, 1, 2]
    assert got == want


def test_is_max():
    input = [5, 5, 1, 2]
    got = is_max(input, 0)
    want = True
    assert got == want


def test_part_one():
    input = """
30373
25512
65332
33549
35390
"""
    got = part_one(input)
    want = 21
    assert got == want


# def test_part_two():
#     want = 1
#     got = 2
#     assert want == got
