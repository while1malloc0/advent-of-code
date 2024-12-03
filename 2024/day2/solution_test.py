from .solution import *

FIXME = None


def test_is_safe_ascending_safely():
    input = [1, 3, 6, 7, 9]
    want = True
    got = is_safe(input)
    assert got == want


def test_is_safe_ascending_unsafely():
    input = [1, 3, 2, 4, 5]
    want = False
    got = is_safe(input)
    assert got == want


def test_is_safe_descending_safely():
    input = [7, 6, 4, 2, 1]
    want = True
    got = is_safe(input)
    assert got == want


def test_is_safe_too_large_interval():
    input = [1, 2, 7, 8, 9]
    want = False
    got = is_safe(input)
    assert got == want


def test_e2e_p1():
    got = p1("example.txt")
    want = 2
    assert got == want


def test_e2e_p2():
    got = p2("example.txt")
    want = 4
    assert want == got
