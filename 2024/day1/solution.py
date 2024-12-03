from collections import Counter
from typing import List


def p1(input: str) -> int:
    with open(input) as content:
        content = content.read()
        pairs = parse_pairs(content)
        sorted_pairs = sort_pairs(pairs)
        distances = distances_from_cols(sorted_pairs)
        result = sum(distances)
        return result


def p2(input: str) -> int:
    with open(input) as f:
        content = f.read()
        pairs = parse_pairs(content)
        similarities = calc_occurances(pairs)
        result = sum(similarities)
        return result


def calc_occurances(pairs: List[List[int]]) -> List[int]:
    lhs = [x[0] for x in pairs]
    rhs = [x[1] for x in pairs]
    c = Counter(rhs)
    return [c[i] * i for i in lhs]


def distances_from_cols(cols: List[List[int]]) -> List[int]:
    distances = []
    for pair in cols:
        distances.append(abs(pair[0] - pair[1]))
    return distances


def sort_pairs(pairs: List[List[int]]) -> List[List[int]]:
    lhs = sorted([x[0] for x in pairs])
    rhs = sorted([x[1] for x in pairs])
    return [[lhs[x], rhs[x]] for x in range(len(pairs))]


def parse_pairs(input: str) -> List[List[int]]:
    num_spaces = 2
    pairs = []
    for row in input.splitlines():
        row = row.strip()
        if row == "":
            continue
        raw_pair = row.split(" ")
        pairs.append([int(raw_pair[0]), int(raw_pair[1 + num_spaces])])
    return pairs


if __name__ == "__main__":
    part1 = p1("input.txt")
    print("Part1: {}".format(part1))
    part2 = p2("input.txt")
    print("Part2: {}".format(part2))
