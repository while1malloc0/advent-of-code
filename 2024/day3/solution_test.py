from .solution import *


FIXME = None


def test_e2e_p1():
    got = p1("example.txt")
    want = 161
    assert got == want


def test_exec_instr():
    input = ["mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"]
    want = [8, 25, 88, 40]
    got = exec_instr(input)
    assert want == got


def test_extract_mults():
    input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
    want = ["mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"]
    got = extract_mults(input)
    assert want == got


def test_extract_mults_and_state():
    input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
    want = ["mul(2,4)", "don't()", "mul(5,5)", "mul(11,8)", "do()", "mul(8,5)"]
    got = extract_mults_and_state(input)
    assert want == got


def test_filter_instrs():
    input = ["mul(2,4)", "don't()", "mul(5,5)", "mul(11,8)", "do()", "mul(8,5)"]
    want = ["mul(2,4)", "mul(8,5)"]
    got = filter_instrs(input)
    assert want == got


def test_e2e_p2():
    got = p2("example.2.txt")
    want = 48
    assert want == got
