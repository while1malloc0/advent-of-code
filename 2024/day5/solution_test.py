from .solution import *
import pytest

FIXME = None

example_rule_table = """
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13
"""


def test_e2e_p1():
    got = p1("example.txt")
    want = 143
    assert got == want


def test_middle_number():
    got = middle_number([75, 47, 61, 53, 29])
    want = 61
    assert want == got


def test_middle_number_short():
    got = middle_number([75, 29, 13])
    want = 29
    assert want == got


def test_rule_check():
    want = True
    got = Rule(75, 29).check([75, 47, 61, 53, 29])
    assert want == got


def test_rules_check():
    subject = Rules.from_str(example_rule_table)
    got = subject.check([75, 47, 61, 53, 29])
    want = True
    assert want == got


def test_rules_check_break():
    subject = Rules.from_str(example_rule_table)
    got = subject.check([97, 13, 75, 29, 47])
    want = False
    assert want == got


def test_rule_from_str():
    input = "47|53"
    want = Rule(47, 53)
    got = Rule.from_str(input)
    assert want == got


def test_rules_correct():
    subject = Rules.from_str(example_rule_table)
    input = [75, 97, 47, 61, 53]
    want = [97, 75, 47, 61, 53]
    got = subject.correct(input)
    assert want == got


def test_rules_correct_short():
    subject = Rules.from_str(example_rule_table)
    input = [61, 13, 29]
    want = [61, 29, 13]
    got = subject.correct(input)
    assert want == got


def test_e2e_p2():
    got = p2("example.txt")
    want = 123
    assert want == got
