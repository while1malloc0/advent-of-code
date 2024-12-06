import pytest
from .solution import *

FIXME = None


def test_e2e_p1():
    got = p1("example.txt")
    want = 18
    assert got == want


def test_find_xmas():
    # fmt: off
    input = [
        ["M","M","M","S","X","X","M","A","S","M"],
        ["M","S","A","M","X","M","S","M","S","A"],
        ["A","M","X","S","X","M","A","A","M","M"],
        ["M","S","A","M","A","S","M","S","M","X"],
        ["X","M","A","S","A","M","X","A","M","M"],
        ["X","X","A","M","M","X","X","A","M","A"],
        ["S","M","S","M","S","A","S","X","S","S"],
        ["S","A","X","A","M","A","S","A","A","A"],
        ["M","A","M","M","M","X","M","M","M","M"],
        ["M","X","M","X","A","X","M","A","S","X"],
    ]
    # fmt: on
    want = 18
    got = find_xmas(input)
    assert want == got


def test_e2e_p2():
    got = p2("example.txt")
    want = 9
    assert want == got
