from typing import List


def p1(input: str) -> int:
    with open(input) as content:
        lines = content.readlines()
        chars = [line.split(" ") for line in lines]
        reports = []
        for c in chars:
            reports.append([int(x) for x in c])
        safe = [report for report in reports if is_safe(report)]
        result = len(safe)
        return result


def is_safe(maybe_safe: List[int]) -> bool:
    ascending = False
    if maybe_safe[0] < maybe_safe[1]:
        ascending = True
    for i in range(len(maybe_safe) - 1):
        if maybe_safe[i] == maybe_safe[i + 1]:
            return False
        if ascending:
            if maybe_safe[i] > maybe_safe[i + 1]:
                return False
        else:
            if maybe_safe[i + 1] > maybe_safe[i]:
                return False
        if abs(maybe_safe[i] - maybe_safe[i + 1]) > 3:
            return False
    return True


def is_sorta_safe(maybe_safe: List[int]) -> bool:
    for i in range(len(maybe_safe)):
        cp = maybe_safe.copy()
        cp.pop(i)
        if is_safe(cp):
            return True
    return False


def p2(input: str) -> int:
    with open(input) as content:
        lines = content.readlines()
        chars = [line.split(" ") for line in lines]
        reports = []
        for c in chars:
            reports.append([int(x) for x in c])
        safe = [report for report in reports if is_sorta_safe(report)]
        result = len(safe)
        return result


if __name__ == "__main__":
    part1 = p1("input.txt")
    print("Part1: {}".format(part1))
    part2 = p2("input.txt")
    print("Part2: {}".format(part2))
