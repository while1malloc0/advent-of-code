from typing import List
import re


def p1(input: str) -> int:
    with open(input) as content:
        mults = extract_mults(content.read())
        products = exec_instr(mults)
        result = sum(products)
        return result


def exec_instr(instrs: List[str]) -> List[int]:
    outs = []
    for inst in instrs:
        inst = inst.replace("mul(", "")
        inst = inst.replace(")", "")
        lhs, rhs = [int(x) for x in inst.split(",")]
        outs.append(lhs * rhs)
    return outs


def extract_mults(input: str) -> List[str]:
    matches = re.findall(r"mul\(\d+,\d+\)", input)
    return matches


def extract_mults_and_state(input: str) -> List[str]:
    matches = re.findall(r"do\(\)|mul\(\d+,\d+\)|don't\(\)", input)
    return matches


def filter_instrs(instrs: List[str]) -> List[str]:
    filtered = []
    active = True
    for instr in instrs:
        if instr == "do()":
            active = True
        elif instr == "don't()":
            active = False
        elif active:
            filtered.append(instr)
        else:  # nothing to do, active is false
            pass
    return filtered


def p2(input: str) -> int:
    with open(input) as content:
        instrs = extract_mults_and_state(content.read())
        instrs = filter_instrs(instrs)
        mults = exec_instr(instrs)
        result = sum(mults)
        return result


if __name__ == "__main__":
    part1 = p1("input.txt")
    print("Part1: {}".format(part1))
    part2 = p2("input.txt")
    print("Part2: {}".format(part2))
